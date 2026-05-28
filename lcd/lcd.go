package lcd

import "fmt"

type Screen struct {
	Width  int
	Height int
	Pixels [][]bool
}

func NewScreen(width, height int) *Screen {
	pixels := make([][]bool, height)

	for y := range pixels {
		pixels[y] = make([]bool, width)
	}

	return &Screen{
		Width:  width,
		Height: height,
		Pixels: pixels,
	}
}

func (s *Screen) SetPixel(x, y int, value bool) {
	if x >= 0 && x < s.Width && y >= 0 && y < s.Height {
		s.Pixels[y][x] = value
	}
}

func (s *Screen) GetPixel(x, y int) bool {
	if x >= 0 && x < s.Width && y >= 0 && y < s.Height {
		return s.Pixels[y][x]
	}
	
	return false
}

func (s *Screen) Clear() {
	for y := range s.Pixels {
		for x := range s.Pixels[y] {
			s.Pixels[x][y] = false
		}
	}
}

func (s *Screen) Render() {
	for y := 0; y < s.Height; y += 2 {
		for x := 0; x < s.Width; x++ {
			top := s.GetPixel(x, y)
			bottom := s.GetPixel(x, y+1)

			if top && bottom {
				fmt.Print("█")
			} else if top {
				fmt.Print("▀")
			} else if bottom {
				fmt.Print("▄")
			} else {
				fmt.Print(" ")
			}
		}
		
		fmt.Println()
	}
}
