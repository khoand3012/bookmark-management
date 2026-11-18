# Bookmark Management

A small HTTP service written in Go, built with the [Gin](https://github.com/gin-gonic/gin) web framework. It is a learning project that follows a layered ("clean") architecture: handlers talk to services through interfaces, and all wiring happens in one place.

Right now it exposes a single endpoint that generates random passwords.

## Requirements

- Go 1.27.3 or newer (see `go.mod`)
- [mockery](https://github.com/vektra/mockery) v2, only if you want to regenerate the test mocks

## Quick start

```bash
git clone https://github.com/khoand3012/bookmark-management.git
cd bookmark-management
go run ./cmd/api
```

The server starts on port `8081`. In another terminal:

```bash
curl http://localhost:8081/genpass
```

```json
{"password":"aB3dE9xYz1"}
```

## Configuration

Configuration is read from environment variables when the app starts (using [envconfig](https://github.com/kelseyhightower/envconfig)). Every variable has a default, so none are required.

| Variable   | Default           | Description                                              |
|------------|-------------------|----------------------------------------------------------|
| `APP_PORT` | `8081`            | Port the HTTP server listens on.                         |
| `APP_NAME` | `bookmark-service`| Name of the application. Not used by any endpoint yet.   |

Example:

```bash
APP_PORT=9000 go run ./cmd/api
```

On Windows PowerShell:

```powershell
$env:APP_PORT = "9000"; go run ./cmd/api
```

## API

### `GET /genpass`

Generates a random 10-character password made of letters (`A-Z`, `a-z`) and digits (`0-9`), using a cryptographically secure random source (`crypto/rand`).

**Success: `200 OK`**

```json
{"password":"aB3dE9xYz1"}
```

**Failure: `500 Internal Server Error`**

```json
{"error":"Failed to generate password"}
```

Any other path or method returns `404 page not found`.

## Project structure

```
cmd/api/                  Entry point: loads config, builds the engine, starts the server
internal/api/             Config, Gin engine and route registration (the only place that wires layers together)
internal/handler/         HTTP layer: reads the request, calls a service, writes the response
internal/service/         Business logic, exposed as interfaces
internal/service/mocks/   Generated mocks of the service interfaces (mockery)
internal/test/
  integration_test/       End-to-end tests that call the real engine through HTTP
```

How a request flows:

```
HTTP request -> Gin route -> handler -> service (interface) -> result -> JSON response
```

Dependencies point inward. Handlers know about service interfaces, services know nothing about HTTP or Gin, and `internal/api` is the only package that builds concrete objects and passes them in through constructors.

## Development

### Run the tests

```bash
go test ./...
```

With coverage:

```bash
go test -cover ./...
```

### Regenerate the mocks

Mocks are generated from the `//go:generate` comments above each service interface. Install mockery v2 once, then run:

```bash
go install github.com/vektra/mockery/v2@v2.53.8
go generate ./internal/service/...
```

Do not edit the files in `internal/service/mocks/` by hand; they are overwritten on every generation.

### Adding a new endpoint

1. Add a service (an interface plus an implementation) in `internal/service/`, with a `//go:generate mockery` comment above the interface.
2. Add a handler in `internal/handler/` that takes the service interface in its constructor.
3. Register the route in `initRoutes` in `internal/api/api.go`.
4. Add a handler unit test (using the generated mock) and an integration test.
