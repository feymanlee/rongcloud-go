# Repository Guidelines

## Project Structure & Module Organization
This repository is a Go SDK for RongCloud IM server APIs. The root entry point is `rongcloud.go`, which exposes lazy-loaded module accessors such as `User()`, `Message()`, and `Callback()`. Feature modules live in top-level directories like `user/`, `group/`, `chatroom_kv/`, and `callback/`.

Most modules follow the same layout:
- `api.go`: path constants, exported `API` interface, and the module implementation
- `types.go`: request and response structs
- `api_test.go`: module tests

Shared internals live under `internal/`: `internal/core/` contains the HTTP client, auth, and failover logic; `internal/testutil/` contains the mock client used by unit tests.

## Build, Test, and Development Commands
Use the standard Go toolchain from the repo root:

- `go mod download`: fetch dependencies
- `go build ./...`: build all packages
- `go vet ./...`: run static analysis
- `go test ./...`: run the full test suite
- `go test ./... -race -coverprofile=coverage.out`: match CI validation
- `go test ./message -run TestContent -v`: run a focused package or test

CI runs build, vet, race tests, and coverage on Go `1.22`, `1.23`, and `1.24`.

## Coding Style & Naming Conventions
Format all Go code with `gofmt`; keep imports grouped by standard library, then external packages. Follow existing module conventions: directory names use underscores for compound modules (`user_profile/`), while package names stay lowercase without underscores (`package userprofile`).

Prefer concise exported method names that match the remote API, `path...` constants for endpoints, and small `api` structs wrapping `core.Client`. Keep request/response types in `types.go`, and embed `types.BaseResp` in API responses where applicable.

## Testing Guidelines
Write table-driven tests in `*_test.go` and use `internal/testutil.NewMockClient()` for module-level coverage. Use `httptest` when validating transport behavior in `internal/core/` or callback handlers. Add tests for new endpoints, error branches, and any change to request encoding (`Post` vs `PostJSON`).

## Commit & Pull Request Guidelines
Recent history uses Conventional Commits, for example `fix(callback): ...`, `feat(callback): ...`, and `refactor(callback): ...`. Keep commits scoped to one module or behavior change.

Pull requests should describe the affected API surface, note any request/response shape changes, and include tests for new or modified behavior. Link the relevant issue when available, and update `README.md` or module docs when public usage changes.
