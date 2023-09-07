# Variables
SERVICE = realtime-leaderboard-api
GOSEC = github.com/securego/gosec/v2/cmd/gosec

default: lint test security coverage

lint:
	@echo "Running lint checks for $(SERVICE)..."
	golangci-lint run

test:
	@echo "Running tests for $(SERVICE)..."

	go test -v -race ./...

coverage:
	@echo "Generating coverage for $(SERVICE)..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.htmls

security:
	@echo "Running security checks for $(SERVICE)..."
	gosec ./...

run:
	@echo "Running $(SERVICE)..."
	go run cmd/server/main.go

.PHONY: lint test security coverage

