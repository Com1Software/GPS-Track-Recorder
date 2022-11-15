package main

import (
	"fmt"
	"runtime"

	term "github.com/nsf/termbox-go"
)

func main() {
	tracktype := 0
	level := 0
	posctl := 0
	pos := ""
	pos1 := ""
	pos2 := ""
	pos3 := ""
	pos4 := ""
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
				fmt.Println(x)
				if level == 0 {
					switch x {
					case "49":
						tracktype = 1
						level = 1
					case "50":
						tracktype = 2
						level = 1

					}
				}
				if level == 1 {
					switch x {
					case "78":
						pos = "N"
					case "110":
						pos = "N"
					case "83":
						pos = "S"
					case "115":
						pos = "S"
					case "69":
						pos = "E"
					case "101":
						pos = "E"
					case "87":
						pos = "W"
					case "119":
						pos = "W"

					}
					switch posctl {
					case 1:
						pos1 = pos
					case 2:
						pos2 = pos
					case 3:
						pos3 = pos
					case 4:
						pos4 = pos

					}
					posctl++
					fmt.Printf("pos1 %s\n", pos1)
					fmt.Printf("pos2 %s\n", pos2)
					fmt.Printf("pos3 %s\n", pos3)
					fmt.Printf("pos4 %s\n", pos4)

				}

			}
		case term.EventError:
			panic(ev.Err)
		}
		if level == 1 {
			if tracktype == 1 {
				fmt.Println("Mapping Straight-Line Track")
				fmt.Println("Select First GPS Point")
				fmt.Println("N - North")
				fmt.Println("S - South")
				fmt.Println("E - East")
				fmt.Println("W - West")
				fmt.Println("Enter Selection:")

			}
			if tracktype == 2 {
				fmt.Println("Mapping Square Four Point Track")
			}
		}
	}
}

func reset() {
	term.Sync()
}
