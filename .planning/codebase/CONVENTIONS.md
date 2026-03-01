# Coding Conventions

**Analysis Date:** 2026-02-27

## Naming Patterns

**Files:**
- Go source files use lowercase with hyphens or underscores: `posts.go`, `postsLibrary.go`, `blog.go`
- Package names match directory names: `package handlers`, `package models`, `package content`
- No file name changes for different purposes (all lowercase, no capitals)

**Functions:**
- Exported functions use PascalCase: `GetAllPosts()`, `GetPost()`, `NewBlogHandler()`
- Unexported functions use camelCase: `renderTemplate()`
- Methods follow receiver pattern: `(h *BlogHandler) Index()`, `(h *BlogHandler) Post()`

**Variables:**
- Short variable names for loop counters: `i`, `j` in sort operations
- Descriptive names for function parameters: `w http.ResponseWriter`, `r *http.Request`, `tmpl string`, `data interface{}`
- Camel case for multi-word variables: `workDir`, `filesDir`, `pathPrefix`, `postsLibrary`

**Types:**
- Struct names use PascalCase: `Post`, `BlogHandler`
- Struct fields use PascalCase: `Title`, `Content`, `ContentPreview`, `Slug`, `CreatedAt`
- Interface types follow receiver conventions from `net/http`

## Code Style

**Formatting:**
- Golang standard formatting (enforced by `go fmt`)
- Indentation: tab-based (Go standard)
- Line length: no explicit limit enforced, but files stay compact (max 85 lines)
- Spacing: standard Go conventions with blank lines between logical sections

**Linting:**
- golangci-lint v1.60 enforced via CI/CD
- Configuration: runs with default rules and 5-minute timeout
- Trigger: automatic on all pull requests to main branch
- Location: `.github/workflows/go-lint.yml`

## Import Organization

**Order:**
1. Standard library imports (first block)
2. Third-party imports (second block)
3. Local package imports (third block)

**Pattern observed in source files:**
```go
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nickfoden/web-log/internal/content"
	"github.com/nickfoden/web-log/internal/handlers"
)
```

**Path Aliases:**
- Not used (standard absolute imports from module root)
- Internal packages referenced with full path: `github.com/nickfoden/web-log/internal/...`

## Error Handling

**Patterns:**
- Explicit error checking: `if err != nil { ... }`
- HTTP errors via `http.Error()` for handler functions
- Fatal errors via `log.Fatal()` for startup failures
- Return empty values on error: `return models.Post{}` when lookup fails
- No error wrapping or context addition observed

**Examples from codebase:**
- `main.go`: `log.Fatal(err)` for startup errors
- `handlers/blog.go`: `http.Error(w, err.Error(), ...)` for handler errors
- `content/posts.go`: Silent failure with empty return on missing post

## Logging

**Framework:** `log` (Go standard library)

**Patterns:**
- Basic logging via `fmt.Printf()` for informational output: `fmt.Printf("Server listening on port %s\n", port)`
- `log.Fatal()` for startup failures only
- No structured logging framework (no slog, no third-party logger)
- Cache control headers set explicitly in templates: `w.Header().Set("Cache-Control", "...")`

## Comments

**When to Comment:**
- Comments used for logical section markers: `// Handlers`, `// Pages`, `// API`
- Comments for non-obvious operations: explaining template parsing, file serving behavior
- Function documentation comments for exported functions: `// FileServer conveniently sets up...`

**JSDoc/TSDoc:**
- Not applicable (Go, not JavaScript/TypeScript)
- Standard Go doc comments used: `// FileServer conveniently sets up a http.FileServer handler to serve`
- Comments precede the item they document

## Function Design

**Size:**
- Small, focused functions: largest is 21 lines (`renderTemplate`)
- Handler methods: 8-20 lines each
- Main entry functions: well-structured with clear sections

**Parameters:**
- Standard HTTP handler signature: `(w http.ResponseWriter, r *http.Request)`
- Constructor functions return initialized types: `NewBlogHandler(posts []models.Post) *BlogHandler`
- Variadic parameters not used; fixed arguments preferred

**Return Values:**
- Single return value for data lookups: `GetPost(slug string) models.Post`
- Implicit error handling via empty/nil returns
- No explicit error returns in public API (errors logged via HTTP)

## Module Design

**Exports:**
- Clear public API via capitalized functions: `GetAllPosts()`, `NewBlogHandler()`, handler methods
- Unexported helper functions for internal use: `renderTemplate()`
- Receiver-based methods for handler operations

**Barrel Files:**
- Not used; each package has single responsibility
- Packages kept minimal: `models/`, `handlers/`, `content/`

## Project Layout

**Package organization:**
- `internal/` directory prevents external imports
- Separation by concern: `handlers/` for HTTP logic, `content/` for data, `models/` for types
- Main application logic in `main.go` at project root

**Entry point:** `main.go` - initializes router, loads content, starts server

---

*Convention analysis: 2026-02-27*
