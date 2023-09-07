package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DB struct {
	db     *sql.DB
	ctx    context.Context
	cancel func()

	DSN           string
	MigrationsDir string
}

func NewDB(dsn string) *DB {
	db := &DB{
		DSN:           dsn,
		MigrationsDir: "./internal/mysql/migrations",
	}

	db.ctx, db.cancel = context.WithCancel(context.Background())
	return db
}

func (db *DB) Open() error {

	if db.DSN == "" {
		return fmt.Errorf("dsn is empty")
	}

	if sqldb, err := sql.Open("mysql", db.DSN); err != nil {
		return err
	} else {
		db.db = sqldb
	}

	driver, err := mysql.WithInstance(db.db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("cannot create migration driver: %w", err)
	}

	if err := db.Migrate(driver); err != nil {
		return fmt.Errorf("cannot migrate db: %w", err)
	}

	return nil
}

func (db *DB) Close() error {
	if db.db != nil {
		return db.db.Close()
	}
	return nil
}

// Migrate applies the database migrations.
func (db *DB) Migrate(driver database.Driver) error {
	if db.MigrationsDir == "" {
		return fmt.Errorf("migrations directory is empty")
	}

	// Construct the migration source URL
	sourceURL := fmt.Sprintf("file://%s", db.MigrationsDir)
	m, err := migrate.NewWithDatabaseInstance(sourceURL, db.DSN, driver)
	if err != nil {
		return err
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
