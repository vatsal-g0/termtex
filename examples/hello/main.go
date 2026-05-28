package main

import (
	"log"

	"termtex/lcd"
	"termtex/textengine"
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