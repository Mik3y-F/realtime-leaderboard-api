package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Mik3y-F/realtime-leaderboard-api/internal/http"
	"github.com/Mik3y-F/realtime-leaderboard-api/internal/mysql"
	"github.com/Mik3y-F/realtime-leaderboard-api/internal/repository"
	"github.com/Mik3y-F/realtime-leaderboard-api/pkg"
)

const (
	HTTPAddress = "HTTP_ADDRESS"
	DOMAIN      = "DOMAIN"

	// DB env settings.
	MYSQL_HOST          = "MYSQL_HOST"
	MYSQL_PORT          = "MYSQL_PORT"
	MYSQL_DATABASE      = "MYSQL_DATABASE"
	MYSQL_USER          = "MYSQL_USER"
	MYSQL_ROOT_PASSWORD = "MYSQL_ROOT_PASSWORD"
	MYSQL_PASSWORD      = "MYSQL_PASSWORD"
)

func main() {
	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	// Instantiate a new type to represent our application.
	// This type lets us shared setup code with our end-to-end tests.
	m := NewMain()

	// Execute program.
	if err := m.Run(ctx); err != nil {
		m.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Wait for CTRL-C.
	<-ctx.Done()

	// Clean up program.
	if err := m.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type Main struct {
	// MySQL database used by MySQL service implementations.
	DB *mysql.DB

	// HTTP server for handling HTTP communication.
	HTTPServer *http.HttpServer

	// Repository implementations.
	PlayerRepository repository.PlayerRepository
	ScoreRepository  repository.ScoreRepository
}

func NewMain() *Main {

	dsn := getDSN()

	return &Main{
		DB:         mysql.NewDB(dsn),
		HTTPServer: http.NewHttpServer(),
	}
}

// Close gracefully stops the program.
func (m *Main) Close() error {
	if m.HTTPServer != nil {
		if err := m.HTTPServer.Close(); err != nil {
			return err
		}
	}
	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Run executes the program. The configuration should already be set up before
// calling this function.
func (m *Main) Run(ctx context.Context) (err error) {
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	// Create repository implementations.
	playerService := mysql.NewPlayerRepository(m.DB)
	scoreService := mysql.NewScoreRepository(m.DB)

	// Attach to Main for use in tests.
	m.PlayerRepository = playerService
	m.ScoreRepository = scoreService

	// Copy repository implementations to the HTTP server.
	m.HTTPServer.PlayerRepository = playerService
	m.HTTPServer.ScoreRepository = scoreService

	// should probably be in a config file
	httpAddress := pkg.GetEnv(HTTPAddress)
	domain := pkg.GetEnv(DOMAIN)

	// Copy configuration settings to the HTTP server.
	m.HTTPServer.Addr = httpAddress
	m.HTTPServer.Domain = domain

	// Start the HTTP server.
	if err := m.HTTPServer.Open(); err != nil {
		return err
	}

	// If TLS enabled, redirect non-TLS connections to TLS.
	if m.HTTPServer.UseTLS() {
		go func() {
			log.Fatal(http.ListenAndServeTLSRedirect(domain))
		}()
	}

	// Enable internal debug endpoints.
	go func() {
		err := http.ListenAndServeDebug()
		if err != nil {
			log.Fatalf("cannot open debug server: %v", err)
		}
	}()

	log.Printf("running: url=%q debug=http://localhost:6060 dsn=%q", m.HTTPServer.URL(), m.DB.DSN)

	return nil
}

func getDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		pkg.MustGetEnv(MYSQL_USER),
		pkg.MustGetEnv(MYSQL_PASSWORD),
		pkg.MustGetEnv(MYSQL_HOST),
		pkg.MustGetEnv(MYSQL_PORT),
		pkg.MustGetEnv(MYSQL_DATABASE),
	)

	return dsn
}
