# ASCII-ART-WEB-EXPORT

## Author
**[Omafu Samuel (somafu]**

## Description

`ascii-art-web-export` extends the **ascii-art-web** project with a file export feature: the ASCII art result generated in-browser can be downloaded as a `.txt` file, using standard HTTP headers to control the browser's download behavior.

Built entirely with the **Go standard library** — no external dependencies.

## Objectives

- Export the ASCII art result in at least one file format (`.txt`).
- Set correct **file permissions** (`0644` — read/write for owner, read-only for others) on the exported file.
- Use `Content-Type`, `Content-Length`, and `Content-Disposition` headers to trigger a proper browser download instead of inline rendering.
- Handle all failure paths (bad input, missing files, I/O errors) with correct HTTP status codes.

## Project Structure

```
ASCII-ART-WEB-EXPORT/
│
├── banners/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
│
├── static/
│   ├── img.png
│   ├── img1.png
│   └── style.css
│
├── template/
│   ├── about.html
│   └── index.html
│
├── Artbuild.go        # ASCII art generation logic
├── handlers.go         # HTTP route handlers (home, ascii-art, download, about)
├── main.go             # Entry point, route registration, server start
├── go.mod
├── Dockerfile
└── README.md
```

## Routes

| Route         | Method | Handler          | Description                                     |
|---------------|--------|------------------|--------------------------------------------------|
| `/`           | GET    | `homeHandler`    | Renders the form; empty `Result` on first load   |
| `/ascii-art`  | POST   | `asciiHandler`   | Reads `text` + `banner` form values, calls `Artbuilder`, re-renders `index.html` with the result |
| `/download`   | GET    | `downloadHandler`| Regenerates the art from query params, writes/reads it as a file, streams it as an attachment |
| `/about`      | GET    | `aboutHandler`   | Renders `about.html`                             |
| `/static/*`   | GET    | `http.FileServer`| Serves `style.css` and images                    |

## Data Flow: `/download`

1. Result page renders a hidden GET form containing the same `text` and `banner` the user submitted to `/ascii-art`.
2. `downloadHandler` reads `text` and `banner` from `r.FormValue`.
3. Validates path (`/download`), method (`GET`), and required fields (`text`, `banner` non-empty) — returns `400`/`404`/`405` on failure.
4. Calls `Artbuilder(text, banner)` to regenerate the ASCII art (no server-side caching of the previous result).
5. `os.WriteFile("Acii-art.txt", []byte(result), 0644)` — creates the file on disk with owner read/write permission.
6. `os.ReadFile("Acii-art.txt")` — reads the bytes back for serving.
7. Sets response headers:
   ```
   Content-Type: text/plain
   Content-Disposition: attachment; filename="Acii-art.txt"
   Content-Length: <len(data)>
   ```
8. `w.Write(data)` streams the bytes to the client.
9. `os.Remove("Acii-art.txt")` deletes the temp file after the response is sent.

`Content-Disposition: attachment` is the header that forces a download prompt rather than inline rendering; `Content-Length` lets the browser show download progress; `Content-Type: text/plain` tells the browser/OS what kind of file it's receiving.

## Error Handling

| Scenario                                   | Status | Route(s)              |
|---------------------------------------------|--------|------------------------|
| Unknown path                                | 404    | all handlers           |
| Wrong HTTP method                           | 405    | `/ascii-art`, `/download` |
| Missing `text` or `banner`                  | 400    | `/download`            |
| No banner selected                          | —      | `/ascii-art` re-renders page with "Select a banner" message |
| `Artbuilder` returns an error (bad banner path/content) | 500 | `/ascii-art`, `/download` |
| `os.WriteFile` / `os.ReadFile` failure       | 500    | `/download`             |

## Running the Project

```bash
go run .
```

Then open your browser at:

```
http://localhost:8000
```

## Usage

1. Enter your text in the textarea.
2. Select a banner style.
3. Click **Generate Art** to see the ASCII art rendered on the page.
4. Click **Download** to export the result as a `.txt` file.

## Allowed Packages

Standard library only: `net/http`, `html/template`, `os`, `strconv`, `strings`, `fmt`.