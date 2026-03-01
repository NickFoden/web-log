---
phase: quick-1
plan: 1
type: execute
wave: 1
depends_on: []
files_modified:
  - internal/content/posts/1.md
  - internal/content/posts/no-bugs.md
  - internal/content/posts/react-native-requires-current-ruby.md
  - internal/content/posts/reduce-the-cost-of-owning-software.md
  - internal/handlers/blog.go
  - go.mod
  - go.sum
autonomous: true
requirements: []
must_haves:
  truths:
    - "All 4 blog posts render correctly in the browser"
    - "The YouTube iframe in reduce-the-cost-of-owning-software renders correctly"
    - "The ordered list with inline code blocks in react-native-requires-current-ruby renders correctly"
    - "HTML post files are deleted after markdown equivalents are confirmed working"
  artifacts:
    - path: "internal/content/posts/1.md"
      provides: "Markdown source for post 1"
    - path: "internal/content/posts/no-bugs.md"
      provides: "Markdown source for no-bugs post"
    - path: "internal/content/posts/react-native-requires-current-ruby.md"
      provides: "Markdown source for react-native post"
    - path: "internal/content/posts/reduce-the-cost-of-owning-software.md"
      provides: "Markdown source for reduce-cost post"
    - path: "internal/handlers/blog.go"
      provides: "Updated handler that reads .md and converts to HTML via goldmark"
  key_links:
    - from: "internal/handlers/blog.go"
      to: "internal/content/posts/*.md"
      via: "os.ReadFile with .md extension"
      pattern: "ReadFile.*\\.md"
    - from: "internal/handlers/blog.go"
      to: "goldmark.Convert"
      via: "markdown-to-HTML conversion before template render"
      pattern: "goldmark\\.Convert"
---

<objective>
Convert the 4 blog post content files from HTML to Markdown, add goldmark as the markdown renderer, and update the Post handler to parse .md files instead of .html files.

Purpose: Markdown is a cleaner authoring format for blog content than raw HTML. This reduces friction when writing new posts.
Output: 4 .md files, updated handler, goldmark dependency added to go.mod.
</objective>

<execution_context>
@/Users/nicholasfoden/.claude/get-shit-done/workflows/execute-plan.md
@/Users/nicholasfoden/.claude/get-shit-done/templates/summary.md
</execution_context>

<context>
@.planning/STATE.md
</context>

<tasks>

<task type="auto">
  <name>Task 1: Convert HTML post files to Markdown</name>
  <files>
    internal/content/posts/1.md
    internal/content/posts/no-bugs.md
    internal/content/posts/react-native-requires-current-ruby.md
    internal/content/posts/reduce-the-cost-of-owning-software.md
  </files>
  <action>
    Create the following 4 Markdown files. Remove HTML wrapper tags (p, ol, li, div, code) and replace with equivalent Markdown syntax. Strip CSS class attributes — they are presentation concerns handled by the stylesheet, not the content.

    **1.md** — three paragraphs of plain text, convert each `<p>` to a paragraph separated by a blank line:

    ```
    Messing around with go and HTMX and html templates to make a simple blog in a way that is new to me. I've gone through a few iterations of blog approaches/sites and maybe this is the one that will stick ¯\_(ツ)_/¯

    So far this post / content files are html file(s), I know markdown seems like a more popular option for dev blogs etc, but I kinda like the idea of writing vanilla html files for content here, especially when llm assisted for styling or formatting.

    I'm betting on leveraging more classic web such as html css files etc, and that we will continue to see less tooling/abstractions for dev experience since devs are now llm assisted. The fastest sites out there are just cached css and html right, why jump through all the hoops when you can generate the styles instantly and then spot check. This post written out and css file generated/styled by hand quickly as you can tell, but I mean for the future !
    ```

    **no-bugs.md** — two paragraphs of plain text:

    ```
    In the current times, we have incredible tooling. And yet over and over for example, I see Typescript apps kicked off without a proper tsconfig, (allowing implicit any,etc) and/or without having git hooks with basic checks in pre-push and cicd/github actions.

    These days it's critical to not skip this and to use all the tooling and configurations available to us. Especially as devs now spin the wheel with ai generated code day in and day out. Gotta have a solid foundation in place with smart defaults to avoid surprises.
    ```

    **react-native-requires-current-ruby.md** — one paragraph then a numbered list. Each list item has inline code using backticks. The original used a `<div class="code-block"><code>...</code></div>` wrapper — replace with a fenced code block (triple backtick) for each command since that is idiomatic Markdown:

    ```
    Hitting issues with React Native setup or errors when running the out of the box starters? You may need to update or check your ruby version. This caught me up the other day on a new machine.

    1. Install Ruby:

       ```
       brew install ruby
       ```

    2. Get the prefix path for Ruby:

       ```
       brew --prefix ruby
       ```

    3. Add Ruby to your PATH:

       ```
       echo 'export PATH="$(brew --prefix ruby)/bin:$PATH"' >> ~/.zshrc
       ```

    4. Source your updated .zshrc file:

       ```
       source ~/.zshrc
       ```

    5. Verify Ruby installation:

       ```
       ruby -v
       ```

    6. Install CocoaPods:

       ```
       gem install cocoapods
       ```
    ```

    **reduce-the-cost-of-owning-software.md** — one blockquote paragraph, one attribution paragraph, then a raw HTML iframe. goldmark will be configured with `WithUnsafe()` (see Task 2) so raw HTML in the .md file passes through unchanged. Write the iframe exactly as it appears in the original HTML file:

    ```
    > "We don't need to get better at authoring software, our challenge, the financial challenge to any business is not I need to write more code faster, It's I need to reduce the cost of owning my software"

    -Ian Cooper - NDC London 2025

    <iframe width="560" height="315" src="https://www.youtube.com/embed/d8NDgwOllaI?si=h8pImr5NdmaJKLca" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
    ```

    Do NOT delete the .html files in this task — deletion happens after manual verification in Task 3.
  </action>
  <verify>ls internal/content/posts/*.md | wc -l (should print 4)</verify>
  <done>4 .md files exist with correct Markdown content matching the original HTML posts</done>
</task>

<task type="auto">
  <name>Task 2: Add goldmark and update Post handler to render Markdown</name>
  <files>
    go.mod
    go.sum
    internal/handlers/blog.go
  </files>
  <action>
    **Step 1 — Add goldmark dependency:**
    Run: `go get github.com/yuin/goldmark`

    This adds goldmark to go.mod and go.sum.

    **Step 2 — Update internal/handlers/blog.go:**

    Add goldmark imports:
    ```go
    import (
        "bytes"
        "fmt"
        "html/template"
        "net/http"
        "os"

        "github.com/go-chi/chi/v5"
        "github.com/nickfoden/web-log/internal/models"
        "github.com/yuin/goldmark"
        "github.com/yuin/goldmark/renderer/html"
    )
    ```

    Replace the Post handler's file-read block with:
    ```go
    func (h *BlogHandler) Post(w http.ResponseWriter, r *http.Request) {
        slug := chi.URLParam(r, "slug")

        for _, post := range h.posts {
            if post.Slug == slug {
                source, err := os.ReadFile(fmt.Sprintf("internal/content/posts/%s.md", post.Slug))
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                md := goldmark.New(
                    goldmark.WithRendererOptions(
                        html.WithUnsafe(),
                    ),
                )

                var buf bytes.Buffer
                if err := md.Convert(source, &buf); err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                renderTemplate(w, "post.html", map[string]interface{}{
                    "Content": template.HTML(buf.String()),
                    "Title":   "Web Log by Nick Foden",
                    "Post":    post,
                })
                return
            }
        }

        http.NotFound(w, r)
    }
    ```

    `html.WithUnsafe()` is required so the raw `<iframe>` in reduce-the-cost-of-owning-software.md passes through to the rendered HTML instead of being stripped.

    After editing, run: `go build ./...` to confirm no compile errors.
  </action>
  <verify>go build ./... exits with code 0</verify>
  <done>go.mod includes goldmark, handler reads .md files, converts to HTML with raw HTML pass-through enabled, and builds cleanly</done>
</task>

<task type="checkpoint:human-verify">
  <name>Task 3: Verify all posts render correctly, then delete HTML files</name>
  <what-built>4 Markdown post files and an updated handler that converts them to HTML at request time using goldmark</what-built>
  <how-to-verify>
    1. Start the server: `go run .` (or however you normally run it)
    2. Visit each post URL and confirm content looks correct:
       - http://localhost:PORT/ — check index lists all 4 posts
       - http://localhost:PORT/blog/1 — confirm three paragraphs render
       - http://localhost:PORT/blog/no-bugs — confirm two paragraphs render
       - http://localhost:PORT/blog/react-native-requires-current-ruby — confirm numbered list with code blocks renders
       - http://localhost:PORT/blog/reduce-the-cost-of-owning-software — confirm quote paragraph, attribution, and YouTube iframe embed renders
    3. If everything looks correct, type "approved" to proceed with HTML file deletion.
    4. If something looks wrong, describe the issue.
  </how-to-verify>
  <resume-signal>Type "approved" to delete the old .html files, or describe the issue to fix first</resume-signal>
</task>

</tasks>

<verification>
After approval in Task 3, delete the 4 old HTML files:
```bash
rm internal/content/posts/1.html
rm internal/content/posts/no-bugs.html
rm internal/content/posts/react-native-requires-current-ruby.html
rm internal/content/posts/reduce-the-cost-of-owning-software.html
```

Then confirm `go build ./...` still passes and all post URLs still serve correctly.
</verification>

<success_criteria>
- 4 .md files exist with correct Markdown content
- go.mod includes github.com/yuin/goldmark
- Post handler reads .md files and uses goldmark with WithUnsafe() to render HTML
- All 4 posts render visually correct in the browser (paragraphs, lists, code blocks, iframe)
- Old .html files deleted
- go build ./... passes
</success_criteria>

<output>
After completion, create `.planning/quick/1-move-blog-posts-from-html-to-markdown-fo/1-SUMMARY.md`
</output>
