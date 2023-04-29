package screen

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

var colors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorBlue,
	termbox.ColorCyan,
	termbox.ColorWhite,
	termbox.ColorYellow,
	termbox.ColorMagenta,
	termbox.ColorLightGray,
	termbox.ColorRed,
}

type gameScreen struct{}

func (g *gameScreen) RenderAsci(board [][]int) {
	fmt.Println("\n==========")
	for _, e := range board {
		for _, num := range e {
			if num > 0 {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func (g *gameScreen) Render(board [][]int) {
	offsety := 3
	offsetx := 3
	cellWidth := 2
	termbox.Clear(termbox.ColorGreen, termbox.ColorGreen)

	for y, e := range board {
		for x, num := range e {
			color := colors[num]
			for i := 0; i < cellWidth; i++ {
				termbox.SetCell(offsetx+cellWidth*x+i, offsety+y, ' ', color, color)
			}
		}
	}

	termbox.Flush()
}

func New() *gameScreen {
	return &gameScreen{}
}
