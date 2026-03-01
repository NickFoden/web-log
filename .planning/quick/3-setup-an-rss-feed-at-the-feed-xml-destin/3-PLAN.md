---
phase: quick
plan: 3
type: execute
wave: 1
depends_on: []
files_modified:
  - internal/handlers/blog.go
  - main.go
autonomous: true
requirements: []

must_haves:
  truths:
    - "GET /feed.xml returns valid RSS 2.0 XML with correct Content-Type"
    - "RSS feed contains all blog posts with title, link, description, and pubDate"
    - "Clicking the existing /feed.xml link in the footer serves the RSS feed"
  artifacts:
    - path: "internal/handlers/blog.go"
      provides: "Feed handler method on BlogHandler"
      contains: "func (h *BlogHandler) Feed"
    - path: "main.go"
      provides: "Route registration for /feed.xml"
      contains: "feed.xml"
  key_links:
    - from: "main.go"
      to: "internal/handlers/blog.go"
      via: "blogHandler.Feed route"
      pattern: 'r\.Get\("/feed\.xml"'
    - from: "internal/handlers/blog.go"
      to: "internal/models/post.go"
      via: "Post struct fields mapped to RSS item fields"
      pattern: "post\\.Title|post\\.Slug|post\\.ContentPreview|post\\.CreatedAt"
---

<objective>
Add an RSS 2.0 feed served at /feed.xml so readers can subscribe to the blog.

Purpose: The footer already links to /feed.xml but the route does not exist yet. This plan creates the handler and route to serve a valid RSS feed generated from the existing posts data.

Output: A working /feed.xml endpoint returning RSS 2.0 XML with all blog posts.
</objective>

<execution_context>
@/Users/nicholasfoden/.claude/get-shit-done/workflows/execute-plan.md
@/Users/nicholasfoden/.claude/get-shit-done/templates/summary.md
</execution_context>

<context>
@internal/handlers/blog.go
@internal/models/post.go
@internal/content/posts.go
@main.go

<interfaces>
<!-- Key types and contracts the executor needs -->

From internal/models/post.go:
```go
type Post struct {
    Title          string
    Content        string
    ContentPreview string
    Slug           string
    CreatedAt      time.Time
}
```

From internal/handlers/blog.go:
```go
type BlogHandler struct {
    posts []models.Post
}
func NewBlogHandler(posts []models.Post) *BlogHandler
```

From main.go:
```go
posts := content.GetAllPosts()           // []models.Post, sorted newest-first
blogHandler := handlers.NewBlogHandler(posts)
r.Get("/about", blogHandler.About)       // pattern for route registration
```
</interfaces>
</context>

<tasks>

<task type="auto">
  <name>Task 1: Add RSS feed handler to BlogHandler</name>
  <files>internal/handlers/blog.go</files>
  <action>
Add a `Feed` method to `BlogHandler` in `internal/handlers/blog.go` that generates and serves an RSS 2.0 XML feed.

Use Go's `encoding/xml` from stdlib (no external library needed). Define RSS structs inline in the handler file (private to the package):

```go
type rssItem struct {
    XMLName     xml.Name `xml:"item"`
    Title       string   `xml:"title"`
    Link        string   `xml:"link"`
    Description string   `xml:"description"`
    PubDate     string   `xml:"pubDate"`
    GUID        string   `xml:"guid"`
}

type rssChannel struct {
    XMLName       xml.Name  `xml:"channel"`
    Title         string    `xml:"title"`
    Link          string    `xml:"link"`
    Description   string    `xml:"description"`
    LastBuildDate string    `xml:"lastBuildDate"`
    Items         []rssItem `xml:"item"`
}

type rssFeed struct {
    XMLName xml.Name   `xml:"rss"`
    Version string     `xml:"version,attr"`
    Channel rssChannel `xml:"channel"`
}
```

The `Feed` method should:
1. Set base URL to `https://nickfoden.com` (the site's domain)
2. Build `rssItem` for each post in `h.posts` (already sorted newest-first):
   - Title: `post.Title`
   - Link: `baseURL + "/posts/" + post.Slug`
   - Description: `post.ContentPreview`
   - PubDate: `post.CreatedAt.Format(time.RFC1123Z)` (RSS requires RFC 822 format)
   - GUID: same as Link
3. Create the rssFeed with channel title "Web Log by Nick Foden", link to baseURL, description "A web log by Nick Foden"
4. Set LastBuildDate from the first post's CreatedAt (newest), or current time if no posts
5. Set `w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")`
6. Write `xml.Header` (`<?xml version="1.0" encoding="UTF-8"?>`) then marshal the feed with `xml.MarshalIndent(feed, "", "  ")`
7. Add `"encoding/xml"` to the import block

Do NOT use any template file for this -- generate XML directly via encoding/xml.
  </action>
  <verify>
    Run `go build ./...` from the project root -- should compile with no errors.
  </verify>
  <done>BlogHandler has a Feed method that generates valid RSS 2.0 XML from the posts slice using encoding/xml</done>
</task>

<task type="auto">
  <name>Task 2: Register /feed.xml route in main.go</name>
  <files>main.go</files>
  <action>
Add a route in main.go to serve the RSS feed. Add this line alongside the existing page routes (after the `/posts/{slug}` line):

```go
r.Get("/feed.xml", blogHandler.Feed)
```

No new imports needed in main.go.
  </action>
  <verify>
    Run `go build ./...` to confirm compilation succeeds, then run `go run main.go &` in background, curl `http://localhost:8080/feed.xml`, confirm it returns XML with Content-Type `application/rss+xml`, and kill the server. Specifically:
    ```
    go build ./... && echo "BUILD OK"
    ```
  </verify>
  <done>GET /feed.xml returns valid RSS 2.0 XML containing all 4 blog posts with correct titles, links, descriptions, and dates. The existing footer link to /feed.xml now resolves.</done>
</task>

</tasks>

<verification>
1. `go build ./...` compiles without errors
2. Start the server and `curl -s -D- http://localhost:8080/feed.xml` shows:
   - HTTP 200 status
   - Content-Type: application/rss+xml; charset=utf-8
   - XML body with `<rss version="2.0">` root element
   - 4 `<item>` elements (one per blog post)
   - Each item has `<title>`, `<link>`, `<description>`, `<pubDate>`, `<guid>`
3. The XML is well-formed (parseable by any RSS reader)
</verification>

<success_criteria>
- /feed.xml serves valid RSS 2.0 XML with all blog posts
- Content-Type header is application/rss+xml
- Each post includes title, permalink, content preview as description, and publish date
- The existing footer "Feed" link now works
- go build succeeds with no errors
</success_criteria>

<output>
After completion, create `.planning/quick/3-setup-an-rss-feed-at-the-feed-xml-destin/3-SUMMARY.md`
</output>
