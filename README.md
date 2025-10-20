# companies

Manage CRU for a single table company with Go Lang and Gin (no API, no React -- just forms)

## Setup

To run this project locally, ensure you have Go  (1.25) installed and run:

```bash
go run main.go
```

## Docker

To build and run the Docker container for production, use the following commands:

```bash
docker build -f Dockerfile.production -t companies-app .
docker run -p 8089:8080 companies-app
```

Then hit `http://localhost:8089/ping` in your browser.

## Run tests

To run test, use the following command:

```bash
go test ./...
```
### Test Coverage

To check test coverage, use the following command:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
This will generate a coverage report and open it in your default web browser.
