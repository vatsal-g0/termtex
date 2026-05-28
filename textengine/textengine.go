package textengine

import (
	"encoding/json"
	"os"

	"github.com/vatsal-g0/termtex/lcd"
)

type Glyph struct {
	Rows  []uint8
	Width int
}

type Font map[rune]Glyph

func rowToBits(row string) uint8 {
	var bits uint8

	for _, c := range row {
		bits <<= 1
		if c == '#' {
			bits |= 1
		}
	}

	return bits
}

func LoadFont(path string) (Font, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var raw map[string][]string

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	font := make(Font)

	for key, bitmap := range raw {
		runes := []rune(key)

		if len(runes) == 0 {
			continue
		}

		glyph := Glyph{
			Rows:  make([]uint8, len(bitmap)),
			Width: len(bitmap[0]),
		}

		for i, row := range bitmap {
			glyph.Rows[i] = rowToBits(row)
		}

		font[runes[0]] = glyph
	}

	return font, nil
}

func DrawGlyph(screen *lcd.Screen, startX, startY int, glyph Glyph) {
	for y, row := range glyph.Rows {
		for x := 0; x < 8; x++ {
			bit := (row >> (7 - x)) & 1

			if bit == 1 {
				screen.SetPixel(startX+x, startY+y, true)
			}
		}
	}
}

func DrawText(screen *lcd.Screen, font Font, startX, startY int, text string) {
	cursorX := startX

	for _, ch := range text {
		glyph, exists := font[ch]

		if !exists {
			continue
		}

		DrawGlyph(screen, cursorX, startY, glyph)

		cursorX += glyph.Width + 1
	}
}
