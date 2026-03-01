---
phase: quick
plan: 2
type: execute
wave: 1
depends_on: []
files_modified:
  - internal/templates/base.html
  - static/css/style.css
autonomous: true
requirements: []

must_haves:
  truths:
    - "Footer displays an About link that navigates to /about"
    - "Footer displays a GitHub link that opens https://github.com/NickFoden in a new tab"
    - "Links are visually distinct from the copyright text"
  artifacts:
    - path: "internal/templates/base.html"
      provides: "Footer with About and GitHub links"
      contains: "/about"
    - path: "static/css/style.css"
      provides: "Footer link styling"
      contains: "footer a"
  key_links:
    - from: "internal/templates/base.html"
      to: "/about"
      via: "anchor href"
      pattern: "href=.*/about"
    - from: "internal/templates/base.html"
      to: "https://github.com/NickFoden"
      via: "anchor href with target=_blank"
      pattern: "href=.*github.com/NickFoden"
---

<objective>
Add About and GitHub links to the site footer in base.html.

Purpose: Give visitors easy navigation to the About page and the project's GitHub profile from every page on the site.
Output: Updated base.html footer with two links, styled consistently with the site.
</objective>

<execution_context>
@/Users/nicholasfoden/.claude/get-shit-done/workflows/execute-plan.md
@/Users/nicholasfoden/.claude/get-shit-done/templates/summary.md
</execution_context>

<context>
@internal/templates/base.html
@static/css/style.css
</context>

<tasks>

<task type="auto">
  <name>Task 1: Add About and GitHub links to footer</name>
  <files>internal/templates/base.html, static/css/style.css</files>
  <action>
    In internal/templates/base.html, update the footer section (lines 36-48). Add a navigation row of links above the existing copyright paragraph. The links should be:
    1. "About" linking to "/about"
    2. "GitHub" linking to "https://github.com/NickFoden" with target="_blank" and rel="noopener noreferrer"

    Structure the footer as:
    ```html
    <footer>
      <nav class="footer_nav">
        <a href="/about">About</a>
        <a href="https://github.com/NickFoden" target="_blank" rel="noopener noreferrer">GitHub</a>
      </nav>
      <p class="font_sixtyfour">
        ... existing copyright content ...
      </p>
    </footer>
    ```

    In static/css/style.css, add styles for the footer navigation after the existing `footer p` rule (after line 21):

    ```css
    footer nav.footer_nav {
      display: flex;
      justify-content: center;
      gap: 24px;
      margin-bottom: 12px;
    }

    footer nav.footer_nav a {
      font-size: 14px;
      text-decoration: none;
    }

    footer nav.footer_nav a:hover {
      text-decoration: underline;
    }
    ```

    Keep the links simple and consistent with the existing minimal site style. Use plain black links (already handled by the global `a` color rule) with underline on hover.
  </action>
  <verify>
    Run `go build ./...` to confirm no build errors. Then run `go run main.go &` and use `curl -s http://localhost:8080 | grep -A5 'footer'` to verify the footer contains both links. Kill the server after verification.
  </verify>
  <done>
    Footer on every page shows "About" link pointing to /about and "GitHub" link pointing to https://github.com/NickFoden opening in a new tab. Links are centered above the copyright line with consistent spacing.
  </done>
</task>

</tasks>

<verification>
- `go build ./...` succeeds
- Footer HTML contains `<a href="/about">About</a>`
- Footer HTML contains `<a href="https://github.com/NickFoden" target="_blank" rel="noopener noreferrer">GitHub</a>`
- CSS file contains `footer_nav` styles
- Links render on all pages (base.html is the shared layout)
</verification>

<success_criteria>
Every page on the site displays footer links to About and GitHub, styled consistently with the minimal site design.
</success_criteria>

<output>
After completion, create `.planning/quick/2-add-about-link-and-github-link-to-the-si/2-SUMMARY.md`
</output>
