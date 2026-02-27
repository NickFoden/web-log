# Testing Patterns

**Analysis Date:** 2026-02-27

## Test Framework

**Runner:**
- Not detected - no test infrastructure in use

**Assertion Library:**
- Not applicable - no testing framework configured

**Run Commands:**
- No test execution commands defined
- No `go test` integration in Makefile or CI/CD pipeline
- Current Makefile only contains: `dev: air` (development server)

## Test File Organization

**Location:**
- No test files present in codebase
- No `*_test.go` files found
- No test directory structure

**Naming:**
- Go standard would be `*_test.go` suffix, but not implemented
- Example: `posts_test.go`, `blog_test.go` (not present)

**Structure:**
```
Current: No test structure
Expected Go pattern:
src/
├── main.go
├── main_test.go
├── internal/
│   ├── handlers/
│   │   ├── blog.go
│   │   └── blog_test.go
│   ├── content/
│   │   ├── posts.go
│   │   └── posts_test.go
│   └── models/
```

## Test Structure

**Suite Organization:**
- Not implemented

**Expected Go testing pattern:**
```golang
func TestGetAllPosts(t *testing.T) {
	posts := GetAllPosts()
	if len(posts) == 0 {
		t.Error("expected posts, got empty slice")
	}
}

func TestNewBlogHandler(t *testing.T) {
	posts := []models.Post{}
	handler := NewBlogHandler(posts)
	if handler == nil {
		t.Error("expected handler, got nil")
	}
}
```

**Patterns:**
- Standard `testing.T` from Go standard library (not configured)
- Table-driven tests recommended for Go: `for _, tt := range tests { ... }`
- No setup/teardown functions currently in use

## Mocking

**Framework:**
- Not in use - no mocking libraries detected
- No testify/mock, gomock, or similar in dependencies

**Patterns:**
- Not applicable - no test infrastructure

**What to Mock (if tests are added):**
- `http.ResponseWriter` for handler testing
- `os.ReadFile` for post content loading
- Template parsing via interface abstraction

**What NOT to Mock:**
- Domain models (`Post`, `BlogHandler`)
- Pure data functions (`GetAllPosts`)
- Time-based operations (unless testing time-specific behavior)

## Fixtures and Factories

**Test Data:**
- Not implemented

**Expected approach for this codebase:**
```golang
// Factory for creating test posts
func createTestPost(slug string) models.Post {
	return models.Post{
		Title:     "Test Post",
		Slug:      slug,
		CreatedAt: time.Now(),
	}
}

// Fixture: sample posts
var testPosts = []models.Post{
	createTestPost("test-1"),
	createTestPost("test-2"),
}
```

**Location:**
- Would be in `testdata/` directory per `.air.toml` configuration
- Or co-located in `*_test.go` files

## Coverage

**Requirements:**
- Not enforced - no coverage thresholds configured
- No coverage reporting in CI/CD pipeline

**View Coverage (if implemented):**
```bash
go test -cover ./...              # View coverage percentage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out  # Generate HTML report
```

## Test Types

**Unit Tests:**
- Not implemented
- Should test individual functions in isolation:
  - `GetAllPosts()` - sorting and data retrieval
  - `GetPost(slug)` - lookup by slug
  - `NewBlogHandler()` - handler initialization

**Integration Tests:**
- Not implemented
- Would test HTTP handlers with mock `http.ResponseWriter`:
  - `Index()` handler template rendering
  - `Post()` handler with file reading
  - `About()`, `Ai()` handlers

**E2E Tests:**
- Not implemented
- Not detected in configuration

## Common Patterns

**Async Testing:**
- Not applicable (synchronous handlers)

**Error Testing:**
- Not implemented
- Example pattern for Go:
```golang
func TestPostHandlerNotFound(t *testing.T) {
	h := NewBlogHandler([]models.Post{})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/posts/notfound", nil)

	h.Post(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}
```

## Recommended Testing Setup

**To add testing to this project:**

1. **Create test files** following Go convention:
   - `main_test.go` - main package tests
   - `internal/handlers/blog_test.go` - handler tests
   - `internal/content/posts_test.go` - content logic tests
   - `internal/models/post_test.go` - model tests

2. **Test handler functions** using `httptest`:
   ```bash
   go test -v ./internal/handlers
   ```

3. **Add to CI/CD** - update `.github/workflows/go-lint.yml` to include:
   ```yaml
   - name: Run tests
     run: go test -v -race -coverprofile=coverage.out ./...
   ```

4. **Testing dependencies needed:**
   - `httptest` (standard library) for HTTP handler testing
   - Optional: `testify/assert` for more readable assertions

---

*Testing analysis: 2026-02-27*
