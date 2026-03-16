# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Test Commands

```bash
go build ./...                          # Build all packages
go vet ./...                            # Static analysis
go test ./...                           # Run all tests
go test ./... -race                     # Run with race detector
go test ./user/... -v                   # Run tests for a single module
go test ./internal/core/... -v -run TestPostFormSignature  # Run a single test
```

## Architecture

This is a Go SDK for the RongCloud IM Server API (~198 endpoints), organized as a modular, interface-based design.

### Core Pattern

Every module follows the same structure:
- `api.go` — path constants, exported `API` interface, unexported `api` struct, `NewAPI(client core.Client) API` constructor
- `types.go` — request/response structs; responses embed `types.BaseResp`

The main entry point `rongcloud.go` exposes an `RC` interface with lazy-loaded module accessors via `sync.Once`. Each accessor (e.g., `User()`, `Group()`) returns a module's `API` interface.

### Dual Content-Type

RongCloud V1 APIs use `application/x-www-form-urlencoded` via `client.Post(path, map[string]string, resp)`. V2 APIs (user_profile, ultragroup_usergroup, some message/push endpoints) use `application/json` via `client.PostJSON(path, body, resp)`. Both methods are on the `core.Client` interface.

### Auth & Failover

Every request gets headers: `App-Key`, `Nonce`, `Timestamp`, `Signature` (SHA1 of appSecret+nonce+timestamp). The client auto-switches between primary and backup domain on network errors or HTTP 500+, with a 30-second cooldown.

### Testing

Tests use `internal/testutil.MockClient` which implements `core.Client`, records all calls (path, params, body), and returns `{"code":200}` by default. Override behavior via `PostFunc`/`PostJSONFunc`. Core client tests use `net/http/httptest` for integration-level validation of auth headers, signatures, and domain failover.

### Adding a New Module

1. Create `module_name/api.go` and `module_name/types.go` (directory uses underscores for compound names, package name uses camelCase: `package modulename`)
2. Define path constants, `API` interface with Chinese comments, `api` struct with `client core.Client`
3. Add accessor method to `RC` interface and `rc` struct in `rongcloud.go` (with `sync.Once` lazy init)
4. Add tests in `module_name/api_test.go` using `testutil.MockClient`
