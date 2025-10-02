
# Gin Golang Boilerplate

Minimal, clean Gin starter with layered structure.

## Run

```bash
APP_ENV=development HTTP_PORT=8080 go run ./cmd/server
```

## Structure
- cmd/server: entrypoint
- internal/*: app modules
- pkg/*: shared packages
- mocks/: mock data
