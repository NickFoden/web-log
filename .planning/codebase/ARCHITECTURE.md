# Architecture

**Analysis Date:** 2026-02-27

## Pattern Overview

**Overall:** Layered MVC pattern with file-based post storage

**Key Characteristics:**
- HTTP request routing via chi router at application layer
- Handler layer manages HTTP request/response cycles
- Content layer manages post metadata and retrieval
- Models layer defines data structures
- Server-side template rendering with Go's `html/template`
- Static file serving for CSS and assets
- Post content stored as individual HTML files in filesystem

## Layers

**Entry Point & Router Layer:**
- Purpose: Bootstrap server, configure routes, establish middleware
- Location: `main.go`
- Contains: Route definitions, static file server setup, middleware configuration (logging, panic recovery)
- Depends on: chi router, content and handler packages
- Used by: HTTP server at startup

**Handler Layer:**
- Purpose: Process HTTP requests, render templates, coordinate between routing and content retrieval
- Location: `internal/handlers/blog.go`
- Contains: `BlogHandler` struct with methods for Index, About, Ai, and Post routes; template rendering function
- Depends on: models, chi for URL parameter extraction
- Used by: Router layer for handling HTTP requests

**Content Layer:**
- Purpose: Aggregate and retrieve post data from in-memory library
- Location: `internal/content/posts.go`, `internal/content/postsLibrary.go`
- Contains: `GetAllPosts()` function that returns sorted posts, `GetPost()` function for retrieval, `PostsLibrary` map storing post metadata
- Depends on: models
- Used by: Handler layer to populate template data

**Models Layer:**
- Purpose: Define core data structures
- Location: `internal/models/post.go`
- Contains: `Post` struct with Title, Content, ContentPreview, Slug, CreatedAt fields
- Depends on: time package for date handling
- Used by: Content and handler layers

**Template Layer:**
- Purpose: Render HTML responses with dynamic content
- Location: `internal/templates/` (base.html, index.html, post.html, about.html, ai.html)
- Contains: HTML templates using Go template syntax with base template inheritance pattern
- Depends on: Handler layer calling `renderTemplate()`
- Used by: Handler layer for response generation

**Static Files Layer:**
- Purpose: Serve CSS stylesheets and favicon assets
- Location: `static/css/`, `static/assets/`
- Contains: CSS files and favicon
- Depends on: FileServer utility function in main.go
- Used by: Browser for styling and icons

## Data Flow

**Post Index Request:**

1. Client requests GET `/`
2. Router matches to `blogHandler.Index`
3. Handler calls `content.GetAllPosts()` which retrieves all posts from `PostsLibrary` and sorts by CreatedAt descending
4. Handler calls `renderTemplate()` with template name "index.html" and data containing sorted posts
5. Template renderer parses base.html and index.html
6. index.html iterates over posts and renders each as list item with title, date, and preview
7. base.html wraps content with HTML structure, includes CSS, and JavaScript (HTMX, Google Analytics)
8. Response sent to client

**Individual Post Request:**

1. Client requests GET `/posts/{slug}`
2. Router extracts slug parameter and passes to `blogHandler.Post`
3. Handler iterates through in-memory posts to find matching slug
4. Handler reads HTML file from `internal/content/posts/{slug}.html` using `os.ReadFile`
5. HTML content is passed as `template.HTML` to template renderer
6. post.html template renders with post metadata and full HTML content
7. Response sent to client

**Current Year API Request:**

1. Client requests GET `/get_current_year` (triggered by HTMX in footer)
2. Inline handler returns current year as string
3. HTMX replaces footer element with year

**State Management:**

- Post metadata state: Loaded once at startup into `PostsLibrary` map (immutable after initialization)
- Post HTML content: Read from filesystem on-demand for each post request
- Handler state: `BlogHandler` struct holds reference to posts list throughout application lifetime
- No session or user state management

## Key Abstractions

**BlogHandler:**
- Purpose: Encapsulate HTTP handler logic and coordinate post data retrieval with templating
- Examples: `internal/handlers/blog.go`
- Pattern: Receiver methods on `BlogHandler` struct implement http.Handler interface signature

**PostsLibrary:**
- Purpose: Central registry of post metadata; provides single source of truth for post catalog
- Examples: `internal/content/postsLibrary.go`
- Pattern: Map-based lookup by slug for O(1) post retrieval

**renderTemplate Function:**
- Purpose: Standardize template loading, parsing, and execution with consistent error handling
- Examples: `internal/handlers/blog.go` lines 17-37
- Pattern: Helper function abstracting Go template API and response header management

**FileServer:**
- Purpose: Serve static files with proper path handling and redirect logic
- Examples: `main.go` lines 57-76
- Pattern: Utility function wrapping chi's file serving with chi router context integration

## Entry Points

**HTTP Server:**
- Location: `main.go` main() function
- Triggers: Application startup
- Responsibilities: Initialize chi router, configure middleware (logging, panic recovery), load posts into memory, register routes, start HTTP listener on configured port (default 8080)

**Route Handlers:**
- GET `/` → `blogHandler.Index` - displays list of all posts
- GET `/posts/{slug}` → `blogHandler.Post` - displays individual post with full content
- GET `/about` → `blogHandler.About` - displays about page
- GET `/ai` → `blogHandler.Ai` - displays AI-related page
- GET `/get_current_year` → inline handler - returns current year for footer
- GET `/static/*` → FileServer - serves CSS, favicon, and other static assets

## Error Handling

**Strategy:** HTTP error responses with status codes

**Patterns:**
- Template parsing errors → http.Error with 500 status
- File read errors (post HTML) → http.Error with 500 status
- Post not found → http.NotFound with 404 status
- Panic recovery → middleware.Recoverer in chi router

## Cross-Cutting Concerns

**Logging:** chi middleware.Logger provides request/response logging; no custom logging

**Validation:** URL slug matching through linear search in handler; minimal validation

**Authentication:** None implemented; site is public

**Caching:** renderTemplate sets Cache-Control headers to prevent caching of dynamic pages

---

*Architecture analysis: 2026-02-27*
