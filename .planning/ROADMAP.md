# Roadmap: Web Log Cleanup

## Overview

Six targeted cleanup tasks applied to a working personal blog. Phase 1 modernizes Go syntax and error handling — low-risk textual changes. Phase 2 refactors request-path behavior: template caching, O(1) post lookup, and path sanitization. The blog runs identically for visitors throughout; these changes reduce technical debt and close a security gap.

## Phases

**Phase Numbering:**
- Integer phases (1, 2, 3): Planned milestone work
- Decimal phases (2.1, 2.2): Urgent insertions (marked with INSERTED)

Decimal phases appear between their surrounding integers in numeric order.

- [ ] **Phase 1: Code Quality** - Modernize Go syntax and improve error handling
- [ ] **Phase 2: Performance & Security** - Cache templates, use O(1) post lookup, sanitize file paths

## Phase Details

### Phase 1: Code Quality
**Goal**: Codebase uses modern Go idioms and handles errors safely
**Depends on**: Nothing (first phase)
**Requirements**: QUAL-01, QUAL-02, QUAL-03
**Success Criteria** (what must be TRUE):
  1. No `interface{}` appears anywhere in the codebase — only `any` is used
  2. All handler data maps use `map[string]any` consistently
  3. HTTP error responses return generic messages; full error details appear only in server logs
**Plans**: TBD

### Phase 2: Performance & Security
**Goal**: Request handling is efficient and input is validated before filesystem access
**Depends on**: Phase 1
**Requirements**: PERF-01, PERF-02, SECU-01
**Success Criteria** (what must be TRUE):
  1. Templates are parsed once at server startup and reused across all requests — no per-request parsing
  2. Post lookups use direct map access by slug — no linear iteration over the posts list
  3. Post file paths are sanitized with `filepath.Clean()` and slug is validated against the known posts map before any filesystem read
**Plans**: TBD

## Progress

**Execution Order:**
Phases execute in numeric order: 1 → 2

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. Code Quality | 0/TBD | Not started | - |
| 2. Performance & Security | 0/TBD | Not started | - |
