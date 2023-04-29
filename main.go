package main

import (
	"math/rand"
	"simpletetris/screen"
	"simpletetris/tetris"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	const animationSpeed = 50 * time.Millisecond

	ticker := time.NewTimer(time.Duration(animationSpeed))

	game := tetris.New()
	scr := screen.New()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowUp:
					game.Rotate()
				case ev.Key == termbox.KeyArrowDown:
					game.SpeedUp()
				case ev.Key == termbox.KeyArrowLeft:
					game.MoveLeft()
				case ev.Key == termbox.KeyArrowRight:
					game.MoveRight()
				case ev.Key == termbox.KeySpace:
					game.Fall()
				case ev.Ch == 'q':
					//quit
					return
				case ev.Ch == 's':
					//start
					game.Start()
				}
			}
		case <-ticker.C:
			scr.Render(game.GetBoard())
			ticker.Reset(animationSpeed)
		case <-game.FallSpeed.C:
			game.GameLoop()
		}
	}

}

// thansk for watching
