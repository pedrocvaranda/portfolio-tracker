# Contributing

Thanks for considering a contribution.

## Local Checks

Run these commands before opening a pull request:

```bash
gofmt -w ./cmd ./portfolio
go test ./...
```

## Development

Start the web UI locally:

```bash
go run ./cmd --mode web --addr :8080
```

Keep changes small and focused. Add or update tests when behavior changes.
