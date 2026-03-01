# Codebase Concerns

**Analysis Date:** 2026-02-27

## Security Concerns

**Path Traversal in Post File Loading:**
- Issue: Post content files are loaded using user-controlled slug in `os.ReadFile()` with string concatenation
- Files: `internal/handlers/blog.go:70`
- Risk: If slug validation is insufficient, malicious input like `../../../etc/passwd` could read arbitrary files
- Current mitigation: Slugs are defined in `PostsLibrary` map, limiting to known posts; however, no explicit path sanitization
- Recommendation: Use `filepath.Clean()` and validate slug against allowed posts list; consider using a whitelist approach (already exists but should be enforced at parse time)

**Template Path Construction:**
- Issue: Template names are concatenated with directory prefix without validation
- Files: `internal/handlers/blog.go:21`
- Risk: If `tmpl` parameter contains path traversal characters, could load unintended templates
- Current mitigation: All callers pass hardcoded template names
- Recommendation: Use a whitelist of allowed template names or validate template parameter against allowed values

**Hardcoded Google Analytics ID:**
- Issue: Google Analytics tracking ID (G-ETK9JGW8YR) is hardcoded in template
- Files: `internal/templates/base.html:17,26`
- Risk: Tracking ID is exposed in source code and git history; complicates changing analytics providers
- Current mitigation: None - ID is public-facing anyway
- Recommendation: Move to environment variable for production consistency and future flexibility

**Template HTML Injection:**
- Issue: Post content is loaded as raw HTML and cast to `template.HTML()` without sanitization
- Files: `internal/handlers/blog.go:76`
- Risk: If post content files contain malicious scripts, they will be executed in browser
- Current mitigation: Post HTML files are developer-controlled
- Recommendation: Consider HTML sanitization library if user-generated content is ever added; add content security policy headers

## Performance Concerns

**Inefficient Post Lookup:**
- Issue: Post lookup uses linear search through all posts
- Files: `internal/handlers/blog.go:68-81`
- Problem: O(n) lookup for every single post request; scales poorly as post count grows
- Current impact: Negligible for small blog (4 posts) but will become issue as blog grows
- Improvement path: Store post slug as key for O(1) lookup instead of iterating through values

**Template Parsing on Every Request:**
- Issue: `template.ParseFiles()` is called for every HTTP request, re-reading and re-parsing template files
- Files: `internal/handlers/blog.go:19-22`
- Problem: Unnecessary disk I/O and parsing overhead on every single request
- Current impact: Measurable latency on each request; no caching
- Improvement path: Parse templates once at startup and cache in `BlogHandler` or use `template.Must()` with global caching

**Inefficient Post Sorting:**
- Issue: Posts are sorted on every application startup but not cached
- Files: `internal/content/posts.go:15-16`
- Problem: If posts list grows large, sorting happens every startup; could use pre-sorted data structure
- Current impact: Minor but compounds with template parsing overhead
- Improvement path: Pre-sort in `PostsLibrary` or cache sorted result in global variable

## Code Quality Concerns

**Deprecated Go Syntax:**
- Issue: Use of `interface{}` instead of `any` keyword
- Files: `internal/handlers/blog.go:17,45,75`
- Problem: Go 1.18+ has `any` as alias for `interface{}`; current code uses outdated syntax
- Impact: Inconsistent with modern Go practices; go.mod specifies 1.22.6
- Fix approach: Replace all `interface{}` with `any` for consistency

**Inconsistent Type Usage in Handler:**
- Issue: Mix of `map[string]interface{}` and `map[string]any` in same file
- Files: `internal/handlers/blog.go:45,53,60,75`
- Problem: Inconsistent style makes code harder to read; lines 53 and 60 use `any`, others use `interface{}`
- Impact: Minor but indicates rushed or multi-author commits without consistency checks
- Fix approach: Standardize to `map[string]any` throughout

**Missing Error Context:**
- Issue: Error messages returned to HTTP clients lack context
- Files: `internal/handlers/blog.go:24-25,33-35,72-73`
- Problem: Generic errors like `http.Error(w, err.Error(), ...)` expose implementation details and don't help debugging
- Impact: Poor observability; hard to debug issues in production
- Fix approach: Log errors with context server-side, return generic messages to client

**Hard Path Dependencies:**
- Issue: Template and post content paths are relative to working directory
- Files: `internal/handlers/blog.go:20-21,70` and `main.go:28`
- Problem: Templates must be at exact relative path; breaks if binary is run from different directory
- Impact: Application is fragile; fails if run from non-project directory
- Fix approach: Use `embed` package to embed static files or resolve paths relative to binary location

**Unused Parameter in Handler Methods:**
- Issue: HTTP handlers take `*http.Request` parameter but often don't use it
- Files: `internal/handlers/blog.go:43,51,58`
- Problem: Index, Ai, and About handlers receive `r *http.Request` but ignore it (code clarity)
- Impact: Minor - technically correct but suggests potential for future bugs or incomplete implementation
- Fix approach: Remove unused parameters where truly not needed or add early exit logic

## Testing Concerns

**No Test Coverage:**
- Issue: No test files found in codebase
- Files: None (complete gap)
- Risk: No automated verification of core functionality; difficult to refactor safely
- Missing coverage:
  - Post retrieval and sorting logic
  - Template rendering error cases
  - File not found handling
  - Path traversal prevention
  - Route matching
- Priority: High
- Recommendation: Add unit tests for `posts.go` content loading, integration tests for handlers, and path validation tests

## Fragile Areas

**Post Slug Registration:**
- Files: `internal/content/postsLibrary.go:9-38`
- Why fragile: Posts must be manually registered in map AND have corresponding HTML file in `internal/content/posts/` directory
- Sync problem: If slug is added to map but HTML file is missing, 500 error occurs; if HTML exists but not in map, content is unreachable
- Safe modification: Always update both the map entry AND create the HTML file; add tests verifying both exist
- Test coverage: No tests verify that all registered posts have corresponding HTML files

**Static File Directory Resolution:**
- Files: `main.go:23-28`
- Why fragile: Uses `os.Getwd()` which depends on working directory at runtime
- Problem: If binary is symlinked or run from different directory, static files fail to load
- Safe modification: Refactor to use embedded files or resolve paths relative to binary location
- Impact: Assets (CSS, favicon) break silently if startup directory is wrong

**Template Hardcoded Paths:**
- Files: `internal/handlers/blog.go:20-21`
- Why fragile: Relative paths to template files must exist exactly where code expects them
- Problem: Templates cached nowhere; re-parsed on every request; path changes break app silently
- Safe modification: Use `embed` package to embed templates at compile time
- Test coverage: No tests verify template files exist and parse correctly

## Scaling Limitations

**Single-Pass Content Loading:**
- Current state: All posts loaded on startup via `content.GetAllPosts()` called once in `main()`
- Limitation: Growing post count increases startup time
- Impact: Binary becomes slower to start; problematic for container deployments with short grace periods
- Scaling path: Defer post loading until first request, or implement lazy loading per post

**No Pagination or Filtering:**
- Issue: Index page renders all posts at once
- Files: `internal/handlers/blog.go:45-48`
- Impact: As post count grows, index page becomes slow to render and slow to load
- Missing: Post filtering by date range, tag, category; pagination controls
- Scaling path: Implement pagination in templates and handler

**No Caching Headers for Content:**
- Issue: Cache-Control set to no-cache on all templates, but post HTML files served with no cache headers
- Files: `internal/handlers/blog.go:28-30` (disables caching) vs `internal/content/posts/*.html` (no etag or versioning)
- Impact: Browser must refetch static post HTML on every load; no CDN caching possible
- Improvement: Set appropriate Cache-Control headers for post content (immutable once published)

## Dependencies at Risk

**Minimal Dependency Footprint (Positive):**
- Current: Only `github.com/go-chi/chi/v5 v5.2.3` as external dependency
- Status: Very recent version (as of 2025) - good security posture
- No: HTML sanitization library; consider `github.com/microcosm-cc/bluemonday` if user-generated HTML is added

**Missing Standard Tooling:**
- No logging framework (using `log.Fatal()` and `fmt.Printf()`)
- No configuration management (hardcoded paths)
- No error tracking/monitoring
- Impact: Difficult to debug production issues; no structured logging

## Deployment Concerns

**Go Version Version Mismatch:**
- Issue: `cloudbuild.yaml` specifies `golang:1.22` but `app.yaml` specifies `runtime: go122` (older format)
- Files: `cloudbuild.yaml:2,10` and `app.yaml:1`
- Risk: App Engine may interpret `go122` as Go 1.22, but inconsistent specifications could cause runtime mismatches
- Recommendation: Update `app.yaml` to use newer runtime identifier format if available, or clarify Go App Engine version

**Old Version Cleanup in CI/CD:**
- Issue: CloudBuild keeps only 10 most recent versions before deleting
- Files: `cloudbuild.yaml:32`
- Risk: If fast deployment cycle, versions may be deleted too quickly for rollback; no alerting if deletion fails
- Impact: Limited rollback window; errors in deletion could silently prevent cleanup
- Recommendation: Monitor version cleanup logs; consider longer retention policy

## Missing Documentation

**No API Documentation:**
- Issue: `/get_current_year` endpoint serves current year but has no documentation
- Files: `main.go:43-46`
- Problem: Purpose is unclear from code alone; HTMX endpoint relies on JavaScript behavior
- Impact: Future maintainers may not understand why this endpoint exists
- Recommendation: Add inline comments explaining the HTMX usage pattern

**No Post Template Schema:**
- Issue: Post HTML files have no documentation about expected format or structure
- Files: `internal/content/posts/*.html`
- Problem: If additional metadata or structure is needed, unclear how to add it
- Impact: Adding features to post rendering is error-prone
- Recommendation: Create example post template with documentation

---

*Concerns audit: 2026-02-27*
