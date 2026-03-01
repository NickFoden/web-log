# Web Log

## What This Is

A personal blog built in Go, deployed on Google App Engine. Serves blog posts, an about page, and an AI page using server-side rendered HTML templates with chi router. Content is stored as static HTML files with metadata registered in an in-memory map.

## Core Value

A clean, maintainable personal blog that follows modern Go conventions and serves content reliably.

## Requirements

### Validated

- ✓ Blog post listing with date-sorted index — existing
- ✓ Individual post rendering from HTML files — existing
- ✓ About page and AI page — existing
- ✓ Static file serving (CSS, favicon) — existing
- ✓ HTMX-powered current year in footer — existing
- ✓ Google App Engine deployment via Cloud Build — existing
- ✓ Chi router with logging and panic recovery middleware — existing

### Active

- [ ] Replace `interface{}` with `any` throughout codebase
- [ ] Standardize `map[string]any` usage consistently
- [ ] Cache parsed templates instead of re-parsing on every request
- [ ] Replace linear post lookup with O(1) slug-based lookup
- [ ] Improve error handling — log details server-side, return generic messages to clients
- [ ] Add `filepath.Clean()` path sanitization for post file loading

### Out of Scope

- Tests — deferred, not part of this cleanup
- Embed package migration — keeping file-based approach for simplicity
- Pagination — blog is small, not needed yet
- New features — this is purely a cleanup/refactor effort
- Logging framework — current approach sufficient for this scope
- Configuration management — hardcoded paths acceptable for App Engine deployment

## Context

- Brownfield project with existing codebase mapped in `.planning/codebase/`
- CONCERNS.md documents specific issues across security, performance, code quality
- Go 1.22.6 with chi v5.2.3 — minimal dependency footprint
- Deployed on Google App Engine Standard with Cloud Build pipeline
- 4 blog posts currently registered in PostsLibrary

## Constraints

- **Stack**: Go 1.22+ with chi router — no framework changes
- **Deployment**: Google App Engine Standard — must remain compatible
- **Scope**: Cleanup and refactor only — no new features
- **File approach**: Keep file-based templates and content — no embed migration

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Keep file-based templates | Simpler for App Engine, works today | — Pending |
| Skip tests this round | Focus purely on code quality cleanup | — Pending |
| Skip embed migration | Current file approach works, reduces scope | — Pending |

---
*Last updated: 2026-02-27 after initialization*
