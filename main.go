package main

import (
	"fmt"
	"runtime"

	term "github.com/nsf/termbox-go"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()
	fmt.Println("GPS-Track-Recorder")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)

Loop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break Loop
			case term.KeySpace:
				reset()
				fmt.Println("Space")
			case term.KeyArrowUp:
				reset()
				fmt.Println("Arrow Up")
			case term.KeyArrowDown:
				reset()
				fmt.Println("Arrow Down")
			case term.KeyArrowLeft:
				reset()
				fmt.Println("Arrow Left")
			case term.KeyArrowRight:
				reset()
				fmt.Println("Arrow Right")

			default:
				reset()
				fmt.Println("Key : ", ev.Ch)

			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}

func reset() {
	term.Sync()
}
