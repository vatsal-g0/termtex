# TermTex

TermTex is a lightweight terminal-based text rendering system written in Go.  
It simulates a simple LCD-style pixel display and renders bitmap fonts onto it.

---

## Features

- Pixel-based LCD framebuffer
- Monochrome rendering (on/off pixels)
- Bitmap font system (JSON-based)
- Glyph rendering engine
- Half-block terminal rendering for better resolution
- Clean separation between LCD and text engine

---

## How it works

Text rendering pipeline:

```

Text → Glyphs → Bitmaps → LCD Screen → Terminal Output

```

Each character is stored as a bitmap glyph and rendered onto a virtual screen buffer.

---

## Project Structure

```

termtex/
├── lcd/              # LCD framebuffer system
├── textengine/       # Font system + text rendering
├── fonts/            # Bitmap font files (JSON)
├── examples/         # Example programs
└── go.mod

````

---

## Font Format

Fonts are stored in JSON as bitmap rows:

```json
{
  "A": [
    " ### ",
    "#   #",
    "#####",
    "#   #",
    "#   #",
    "#   #",
    "#   #"
  ]
}
````

Each row is converted internally into a bitmask for efficient rendering.

**Note:** The default font file `fonts/default.json` only contains uppercase letters (`A-Z`) and digits (`0-9`).

---

## Running the project

From the project root:

```bash
go run ./examples/hello
```

---

## Build

```bash
go build ./...
```

---

## Notes

* This is not a full terminal UI framework
* Rendering is monochrome (black/white)
* Designed for simplicity and experimentation

---

## License

MIT
