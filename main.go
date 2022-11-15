package main

import (
	"fmt"
	"runtime"

	term "github.com/nsf/termbox-go"
)

func main() {
	tracktype := 0
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()
	fmt.Println("GPS-Track-Recorder")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	fmt.Println("Select Track Type")
	fmt.Println("1 - Straight-Line Track")
	fmt.Println("2 - Square Four Point Track")
	fmt.Println("Enter Selection:")

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
				x := fmt.Sprint(ev.Ch)
				switch x {
				case "49":
					tracktype = 1
				case "50":
					tracktype = 2

				}

			}
		case term.EventError:
			panic(ev.Err)
		}
		if tracktype == 1 {
			fmt.Println("Mapping Straight-Line Track")
		}
		if tracktype == 2 {
			fmt.Println("Mapping Square Four Point Track")
		}
	}
}

func reset() {
	term.Sync()
}
