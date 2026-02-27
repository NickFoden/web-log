# External Integrations

**Analysis Date:** 2026-02-27

## APIs & External Services

**Current Year Endpoint:**
- Internal API endpoint `/get_current_year` returns current year as JSON
  - Implementation: `main.go` lines 43-46
  - No external service dependency - uses Go standard library `time.Now().Year()`

## Data Storage

**Databases:**
- Not used - Application uses in-memory data structures
- Post data is stored in Go source code (`internal/content/postsLibrary.go`)
- No database client or ORM configured

**File Storage:**
- Local filesystem only
- Post content stored as HTML files in `internal/content/posts/`
- Static assets served from `static/` directory via FileServer in `main.go`
- Static file serving: `main.go` lines 28-29 and 59-76

**Caching:**
- None configured - Uses HTTP cache-control headers for browser caching
- Cache-Control header set to `no-cache, no-store, must-revalidate` for dynamic templates
  - Implementation: `internal/handlers/blog.go` lines 28-30

## Authentication & Identity

**Auth Provider:**
- None - Application is public, no authentication required
- No user management or identity system

## Monitoring & Observability

**Error Tracking:**
- None detected - No external error tracking service configured

**Logs:**
- Standard output only
- chi middleware.Logger middleware logs HTTP requests (configured in `main.go` line 20)
- Error logging via Go standard `log` package

## CI/CD & Deployment

**Hosting:**
- Google Cloud Platform - Google App Engine Standard
- Service name: `web-log` (configured in `app.yaml`)
- Automatic HTTPS enforcement enabled

**CI Pipeline:**
- Google Cloud Build (`cloudbuild.yaml`)
  - Step 1: Dependency fetch with golang:1.22 image
  - Step 2: Go binary compilation to `main` executable
  - Step 3: App Engine deployment via gcloud CLI
  - Step 4: Old version cleanup and deletion (keeps 10 most recent versions)
  - Build timeout: 1600 seconds
  - Logging: Cloud Logging only

**Dependency Management:**
- Dependabot configured for Go modules (`/.github/dependabot.yml`)
  - Package ecosystem: gomod
  - Update schedule: weekly
  - Monitors `go.mod` for dependency updates

## Environment Configuration

**Required env vars:**
- PORT - Server port (optional, defaults to 8080)

**Secrets location:**
- No secrets required for operation
- Cloud Build uses `.env` files (if needed) not present in repository
- All configuration in code or environment variables

## Webhooks & Callbacks

**Incoming:**
- None detected

**Outgoing:**
- None detected

---

*Integration audit: 2026-02-27*
