package main

import (
	"log"

	"github.com/vatsal-g0/termtex/lcd"
	"github.com/vatsal-g0/termtex/textengine"
)

func main() {
	screen := lcd.NewScreen(128, 32)

	font, err := textengine.LoadFont(
		"fonts/default.json",
	)

	if err != nil {
		log.Fatal(err)
	}

	textengine.DrawText(
		screen,
		font,
		0,
		0,
		"TERMTEX",
	)

	screen.Render()
}