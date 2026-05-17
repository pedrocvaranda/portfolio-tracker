# Contributing

**Thanks for considering a contribution to Portfolio Tracker!**

---

## Local Checks

Run these commands before opening a pull request:

```bash
gofmt -w ./cmd ./portfolio
go test ./...
```

---

## Development

Start the web UI locally:

```bash
go run ./cmd --mode web --addr :8080
```

---

## Guidelines

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Keep changes small and focused. Add or update tests when behavior changes.
