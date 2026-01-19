# Repository Guidelines

## Project Structure & Module Organization
- `main.go` is the application entry point.
- `controllers/`, `routers/`, and `middlewares/` handle HTTP routing and request flow.
- `models/` contains data models; `db/` is used for the SQLite database file.
- `feature/` holds trading features; `feature/strategy/` contains strategy logic.
- `views/` and `static/` serve the web UI; `lang/` contains i18n resources.
- `conf/` stores configuration templates; `notify/` handles notifications.
- `tests/` contains Beego integration tests.

## Build, Test, and Development Commands
- `cp conf/app.conf.example conf/app.conf` to create a local config.
- `go install github.com/beego/bee/v2@latest` installs the Beego CLI (ensure `GOPATH/bin` is on `PATH`).
- `go mod tidy` installs module dependencies.
- `bee run` starts the dev server (UI at `http://localhost:3333/zmkm/index.html`).
- `go test ./...` runs all Go tests.
- `bee pack -be GOOS=windows` builds a Windows package.

## Coding Style & Naming Conventions
- Go code should be formatted with `gofmt` (tabs for indentation).
- Keep package names lowercase and align file naming with existing patterns (for example, `controllers/feature.go`, `feature/feature_notice.go`).
- Prefer small, focused functions and reuse existing helpers in `utils/`.

## Testing Guidelines
- Use GoConvey for new tests to match `tests/default_test.go`.
- Name test files with `*_test.go` and use `Convey` blocks for assertions.
- Run targeted tests with `go test ./tests -run TestBeego` when validating HTTP endpoints.

## Commit & Pull Request Guidelines
- Commit messages in history use short prefixes like `feat:`, `fix:`, and `doc:`; follow that pattern when possible.
- PRs should include a concise summary, testing notes, and any config changes.
- Link relevant issues and add UI screenshots when views or static assets change.

## Configuration & Security Tips
- Keep secrets out of Git; copy `conf/app.conf.example` to `conf/app.conf` locally.
- If you use SQLite, the DB lives under `db/`; for MySQL, update `conf/app.conf` accordingly.
