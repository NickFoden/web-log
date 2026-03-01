---
phase: quick-1
plan: 1
status: complete
started: 2026-03-01
completed: 2026-03-01
---

## Summary

Migrated all 4 blog post content files from HTML to Markdown and added goldmark as the markdown rendering engine.

## Tasks Completed

| # | Task | Commit | Status |
|---|------|--------|--------|
| 1 | Convert HTML post files to Markdown | 995e868 | Complete |
| 2 | Add goldmark and update Post handler | 7568b90 | Complete |
| 3 | Verify rendering and delete HTML files | c5ec251 | Complete (human-verified) |

## Key Changes

- **4 Markdown files created:** `internal/content/posts/{1,no-bugs,react-native-requires-current-ruby,reduce-the-cost-of-owning-software}.md`
- **Handler updated:** `internal/handlers/blog.go` — `Post` handler reads `.md` files and converts to HTML via goldmark with `html.WithUnsafe()` (required for YouTube iframe passthrough)
- **Dependency added:** `github.com/yuin/goldmark v1.7.16`
- **4 HTML files deleted** after human verification confirmed all posts render correctly

## Deviations

None.

## Self-Check: PASSED

- [x] All 4 .md files exist with correct content
- [x] goldmark in go.mod
- [x] Handler reads .md and converts with WithUnsafe()
- [x] go build ./... passes
- [x] Old .html files deleted
- [x] Human verified rendering
