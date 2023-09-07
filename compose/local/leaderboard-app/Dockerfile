# Start from the official Go image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Install git as it's commonly required for fetching Go dependencies
RUN apt-get update && apt-get install -y git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install golang-migrate
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy the entire content. This step isn't entirely necessary since you're mounting your local volume,
# but it ensures the image has the code.
COPY . .

# Expose the port the app runs on
EXPOSE 8080

# This command starts your application
CMD ["go", "run", "./cmd/server/main.go"]
