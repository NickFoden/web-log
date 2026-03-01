# Requirements: Web Log Cleanup

**Defined:** 2026-02-27
**Core Value:** A clean, maintainable personal blog that follows modern Go conventions and serves content reliably

## v1 Requirements

Requirements for this cleanup round. Each maps to roadmap phases.

### Code Quality

- [ ] **QUAL-01**: All `interface{}` replaced with `any` keyword throughout codebase
- [ ] **QUAL-02**: Consistent `map[string]any` usage in all handler data maps
- [ ] **QUAL-03**: Error messages log details server-side and return generic messages to HTTP clients

### Performance

- [ ] **PERF-01**: Templates parsed once at startup and cached in BlogHandler, not re-parsed per request
- [ ] **PERF-02**: Post lookup uses O(1) slug-based map access instead of linear iteration

### Security

- [ ] **SECU-01**: Post file loading uses `filepath.Clean()` and validates slug against known posts map before filesystem access

## v2 Requirements

Deferred to future work. Tracked but not in current roadmap.

### Testing

- **TEST-01**: Unit tests for content loading and post retrieval
- **TEST-02**: Integration tests for HTTP handlers
- **TEST-03**: Path traversal prevention tests

### Infrastructure

- **INFR-01**: Migrate to Go embed package for templates and static files
- **INFR-02**: Add structured logging framework
- **INFR-03**: Move Google Analytics ID to environment variable
- **INFR-04**: Add pagination for post index

## Out of Scope

| Feature | Reason |
|---------|--------|
| New features | This is purely cleanup/refactor |
| Embed migration | Keeping file-based approach per user preference |
| Test coverage | Deferred to future round |
| Logging framework | Current approach sufficient |
| Configuration management | Hardcoded paths work for App Engine |

## Traceability

Which phases cover which requirements. Updated during roadmap creation.

| Requirement | Phase | Status |
|-------------|-------|--------|
| QUAL-01 | Phase 1 | Pending |
| QUAL-02 | Phase 1 | Pending |
| QUAL-03 | Phase 1 | Pending |
| PERF-01 | Phase 2 | Pending |
| PERF-02 | Phase 2 | Pending |
| SECU-01 | Phase 2 | Pending |

**Coverage:**
- v1 requirements: 6 total
- Mapped to phases: 6
- Unmapped: 0

---
*Requirements defined: 2026-02-27*
*Last updated: 2026-02-27 after roadmap creation*
