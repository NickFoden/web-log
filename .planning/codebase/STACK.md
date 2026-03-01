# Technology Stack

**Analysis Date:** 2026-02-27

## Languages

**Primary:**
- Go 1.22.6 - Complete application implementation, HTTP handlers, content management, and routing

## Runtime

**Environment:**
- Go 1.22.6 (specified in `go.mod`)
- Google App Engine Standard (Go 1.22 runtime as configured in `app.yaml`)

**Package Manager:**
- Go Modules - Standard Go dependency management
- Lockfile: Present (`go.sum` exists)

## Frameworks

**Core:**
- chi/v5 v5.2.3 - HTTP routing and middleware framework (`github.com/go-chi/chi/v5`)
  - Middleware components: Logger middleware, Recoverer middleware
  - Used in `main.go` for route setup and request handling

**Build/Dev:**
- air - Live reload development tool (specified in `Makefile` and `.air.toml`)
  - Configuration: `.air.toml` defines build settings
  - Binary output: `./tmp/main`
  - Watches: `.go`, `.tpl`, `.tmpl`, `.html` files
  - Excludes: `assets`, `tmp`, `vendor`, `testdata` directories

**Linting:**
- golangci-lint v1.60 - Code quality and lint checking
  - Configuration: GitHub Actions workflow `.github/workflows/go-lint.yml`
  - Timeout: 5 minutes

## Key Dependencies

**Critical:**
- github.com/go-chi/chi/v5 v5.2.3 - HTTP router and middleware framework
  - Provides request routing, middleware support, and URL parameter handling
  - Essential for the web server operation

**Standard Library Usage:**
- `net/http` - HTTP server and request/response handling
- `html/template` - HTML template rendering and parsing
- `os` - File system access, environment variables, working directory operations
- `path/filepath` - File path manipulation
- `strings` - String utilities for route pattern validation
- `time` - Time operations (current year calculation, post creation timestamps)
- `fmt` - Formatted I/O and logging
- `log` - Application logging

## Configuration

**Environment:**
- PORT environment variable - Server port (defaults to 8080 if not set)
- GOPATH - Go workspace path (set in Cloud Build)
- GO_VERSION - Go version specification (1.22 in Cloud Build)

**Build:**
- Build configuration: `.air.toml` - Development hot-reload settings
- Cloud Build configuration: `cloudbuild.yaml` - Google Cloud deployment pipeline
- App Engine configuration: `app.yaml` - GAE runtime and routing rules

## Platform Requirements

**Development:**
- Go 1.22.6 or compatible
- air tool for hot-reloading during development
- Standard Go toolchain (`go build`, `go get`)

**Production:**
- Google App Engine Standard (Go 1.22 runtime)
- Automatic HTTPS enforcement (secure: always in `app.yaml`)

---

*Stack analysis: 2026-02-27*
