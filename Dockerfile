# Use official Golang image as base
FROM golang:1.24-alpine

# Install swag (swagger generator)
RUN go install github.com/swaggo/swag/cmd/swag@latest


# Create working directory
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Install make and other essential build tools
RUN apk update && apk add --no-cache make

# Run make to generate docs and build the app
RUN make build

# Expose port (adjust if needed)
EXPOSE 8080

# Run the built binary
CMD ["./bin/IeltS"]
