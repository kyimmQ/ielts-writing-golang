# IELTS Writing API

A RESTful API service built with Go (Golang) for IELTS writing assessment and management.

## Technologies

- Go 1.24
- Gin Web Framework
- MongoDB
- Swagger (swag)
- Docker & Docker Compose

## Prerequisites

- Go 1.24 or higher
- MongoDB
- Docker (optional)
- Make

## Configuration

1. Copy the sample configuration file:

```bash
cp configs/sample.yaml configs/env.yaml
```

2. Update the `configs/env.yaml` with your settings:

```yaml
server:
  port: 8080
mongodb:
  uri: your_mongodb_connection_string
jwt:
  secret_key: your_secret_key
  expiry: 3600
```

## Running the Application

### Local Development

1. Install dependencies:

```bash
go mod download
```

2. Generate Swagger documentation:

```bash
make swag
```

3. Build and run:

```bash
make run
```

### Using Docker

1. Build and run using Docker Compose:

```bash
docker-compose up --build
```

## API Documentation

Once the application is running, you can access the Swagger documentation at:

```
http://localhost:8080/swagger/*
```

## Project Structure

```
.
├── cmd/                    # Application entry point
├── configs/               # Configuration files
├── docs/                  # Swagger documentation
├── global/                # Global variables and constants
├── internal/              # Internal application code
│   ├── entity/           # Domain entities
│   ├── initialize/       # Application initialization
│   ├── middlewares/      # HTTP middlewares
│   ├── modules/          # Business logic modules
│   ├── routes/           # HTTP route definitions
│   └── server/           # HTTP server setup
├── pkg/                   # Reusable packages
│   ├── error/            # Error handling
│   ├── hash/             # Password hashing
│   ├── jwt/              # JWT utilities
│   ├── logger/           # Logging utilities
│   ├── response/         # HTTP response helpers
│   ├── settings/         # Configuration settings
│   └── utils/            # Utility functions
└── docker-compose.yml    # Docker compose configuration
```

## Authentication

The API uses JWT tokens for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your_token>
```

## License

[MIT License](LICENSE)
