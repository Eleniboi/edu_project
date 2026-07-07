# ASCII-ART-STYLIZE

## Author
**[Omafu Samuel (Somafu)]**

## Description

`ascii-art-stylize` builds on **ascii-art-web**, focusing purely on the front-end: making the site more appealing, interactive, and user-friendly through CSS — without changing the underlying Go logic.

## Objectives

- Improve visual appeal, interactivity, and intuitiveness of the site.
- Improve usability and give clearer feedback to the user (e.g. on invalid input, loading, empty results).
- Ensure text remains readable against any background/color choice used.
- Keep the design consistent, responsive, and interactive across pages and screen sizes.

## Requirements

- CSS must be present and linked to the HTML.
- Only the Go standard library is used server-side — no external frameworks.
- Follows the project's [good practices](https://acad.learn2earn.ng/api/content/root/01-edu_module/content/good-practices/README.md) guidelines.

## What Changed

- Refined layout, spacing, and typography in `static/style.css`.
- Consistent color scheme and contrast across all pages (home, result, about).
- Responsive layout for smaller screen sizes.
- Clearer visual feedback for actions (e.g. button states, invalid banner selection, result display).

## Running the Project

```bash
go run .
```

Then open:

```
http://localhost:8000
```

## Learn More

- [Principles of good web design](https://acad.learn2earn.ng/api/content/root/01-edu_module/content/README.md)
- [Responsive design](https://smallbiztrends.com/2013/05/what-is-responsive-web-design.html)
- [Interactive design](https://en.m.wikipedia.org/wiki/Interactive_design)
- [Why consistency matters in web design](https://digitalcommunications.wp.st-andrews.ac.uk/2016/04/07/why-is-consistency-important-in-web-design/)