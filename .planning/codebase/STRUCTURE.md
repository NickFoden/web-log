# Codebase Structure

**Analysis Date:** 2026-02-27

## Directory Layout

```
web-log/
├── main.go                          # Application entry point, router config, file server
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── Makefile                         # Build commands
├── README.md                        # Project documentation
├── app.yaml                         # Google Cloud deployment config
├── cloudbuild.yaml                 # Cloud Build pipeline config
├── .air.toml                        # Live reload configuration
├── .gitignore                       # Git ignore patterns
├── .github/                         # GitHub Actions workflows
├── .planning/                       # GSD planning documents
├── .claude/                         # Claude-specific config
├── google-domains/                  # Google Domains configuration files
├── tmp/                             # Temporary build/runtime directory
├── internal/                        # Private application packages
│   ├── handlers/
│   │   └── blog.go                 # HTTP handlers for routes
│   ├── content/
│   │   ├── posts.go                # Post retrieval and sorting logic
│   │   ├── postsLibrary.go         # Post metadata registry
│   │   └── posts/                  # Post HTML content files
│   │       ├── 1.html
│   │       ├── no-bugs.html
│   │       ├── react-native-requires-current-ruby.html
│   │       └── reduce-the-cost-of-owning-software.html
│   ├── models/
│   │   └── post.go                 # Post data structure
│   └── templates/                  # HTML templates
│       ├── base.html               # Base layout template
│       ├── index.html              # Posts list page
│       ├── post.html               # Individual post page
│       ├── about.html              # About page
│       └── ai.html                 # AI page
└── static/                          # Publicly served static files
    ├── css/                         # Stylesheets
    │   └── style.css
    └── assets/                      # Images, favicon, etc
        └── favicon.ico
```

## Directory Purposes

**internal/:**
- Purpose: Go-enforced private package directory; contains all application code
- Contains: Go packages for handlers, content management, models, and templates
- Key files: All .go files and HTML template files

**internal/handlers/:**
- Purpose: HTTP request handlers implementing route logic
- Contains: Handler functions and rendering utilities
- Key files: `blog.go`

**internal/content/:**
- Purpose: Post data management and aggregation
- Contains: Post metadata registry, retrieval functions, HTML content files
- Key files: `posts.go` (retrieval logic), `postsLibrary.go` (metadata), `posts/` subdirectory (content)

**internal/models/:**
- Purpose: Define core domain data structures
- Contains: Post struct definition
- Key files: `post.go`

**internal/templates/:**
- Purpose: HTML templates for rendering responses
- Contains: Base layout and page-specific templates using Go template syntax
- Key files: `base.html` (layout), `index.html` (post list), `post.html` (single post), `about.html`, `ai.html`

**static/:**
- Purpose: Serve public CSS and asset files to clients
- Contains: CSS stylesheets and favicon
- Key files: `css/style.css`, `assets/favicon.ico`

**static/css/:**
- Purpose: Application stylesheets
- Key files: `style.css`

**static/assets/:**
- Purpose: Static assets like favicon
- Key files: `favicon.ico`

## Key File Locations

**Entry Points:**
- `main.go`: Application startup, chi router initialization, static file server setup, port configuration (default 8080), route definitions

**Configuration:**
- `go.mod`: Module declaration and direct dependencies
- `.air.toml`: Live reload configuration for development
- `app.yaml`: Google Cloud App Engine deployment manifest
- `cloudbuild.yaml`: GCP Cloud Build pipeline steps

**Core Logic:**
- `internal/handlers/blog.go`: HTTP handler methods (Index, Post, About, Ai) and renderTemplate utility
- `internal/content/posts.go`: GetAllPosts() for sorted retrieval, GetPost() for single post lookup
- `internal/content/postsLibrary.go`: PostsLibrary map with all post metadata

**Data Models:**
- `internal/models/post.go`: Post struct with Title, Content, ContentPreview, Slug, CreatedAt fields

**Presentation:**
- `internal/templates/base.html`: Master template with header, footer, Google Analytics, HTMX, font links
- `internal/templates/index.html`: Lists all posts with title, date, and preview
- `internal/templates/post.html`: Renders individual post with metadata and HTML content
- `internal/templates/about.html`: About page template
- `internal/templates/ai.html`: AI page template

**Styling:**
- `static/css/style.css`: All application CSS

**Testing:**
- No test files present (See CONCERNS.md)

## Naming Conventions

**Files:**
- Go source files: lowercase with underscores (e.g., `posts.go`, `postsLibrary.go`)
- HTML files: lowercase with hyphens (e.g., `base.html`, `index.html`)
- Static files: lowercase (e.g., `style.css`, `favicon.ico`)

**Directories:**
- Package directories: lowercase plural for collections (e.g., `handlers`, `models`, `templates`)
- Single logical unit: lowercase singular (e.g., `content`)

**Go Packages:**
- Use standard Go package naming (lowercase, no underscores)
- Public types: PascalCase (e.g., `BlogHandler`, `Post`)
- Private functions: camelCase (e.g., `renderTemplate`, `GetAllPosts`)

**Routes:**
- Lowercase paths: `/`, `/about`, `/ai`, `/posts/{slug}`, `/get_current_year`, `/static/*`
- URL parameters: curly braces via chi (e.g., `{slug}`)

**Post Slugs:**
- Lowercase with hyphens: `reduce-the-cost-of-owning-software`, `react-native-requires-current-ruby`, `no-bugs`
- Single number also used: `1`

## Where to Add New Code

**New Feature (Route):**
- Add route registration in `main.go` between lines 37-40
- Add handler method to `BlogHandler` struct in `internal/handlers/blog.go`
- Add HTML template to `internal/templates/{page-name}.html`
- Update `renderTemplate()` calls to include necessary data

**New Page Content:**
- Create HTML template file in `internal/templates/{page-name}.html`
- Follow base.html template inheritance pattern (see index.html for example)
- Use `{{template "base.html" .}}` and `{{ define "content" }}...{{ end }}` structure

**New Post:**
- Add entry to `PostsLibrary` map in `internal/content/postsLibrary.go` with slug key
- Create HTML content file in `internal/content/posts/{slug}.html`
- Entry must include: Title, ContentPreview, Slug, CreatedAt (time.Date format)

**New Utilities:**
- Shared helpers go in new package under `internal/` (e.g., `internal/utils/`)
- Follow Go package conventions: lowercase package name, exported functions PascalCase
- Import in relevant files (handlers or content packages)

**Styling:**
- Add CSS classes to `static/css/style.css`
- Use class names referenced in templates and HTML content files
- Font families defined in base.html includes (Roboto Condensed, Sixtyfour)

**Static Assets:**
- Place images/icons in `static/assets/`
- Reference in templates as `/static/assets/{filename}`
- Update FileServer in main.go if new asset types need special handling

## Special Directories

**tmp/:**
- Purpose: Temporary build artifacts and runtime files
- Generated: Yes (during build/deployment)
- Committed: No (.gitignore)

**google-domains/:**
- Purpose: Domain configuration files for Google Domains
- Generated: No (manually maintained)
- Committed: Yes

**.planning/codebase/:**
- Purpose: GSD codebase analysis documents (ARCHITECTURE.md, STRUCTURE.md, etc.)
- Generated: Yes (by GSD commands)
- Committed: Yes

---

*Structure analysis: 2026-02-27*
