package main

import (
	"fmt"
	"runtime"
      gps "github.com/Com1Software/GPS-Package"
	term "github.com/nsf/termbox-go"
)

func main() {
	port := gps.getSerialPort()
	fmt.Println(port)
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
				fmt.Println(level)
				switch level {
				case 0:
					switch x {
					case "49":
						tracktype = 1
						level = 1
					case "50":
						tracktype = 2
						level = 1

					}

				case 4:
					switch x {
					case "89":
						pos = "N"
					case "121":
						pos = "N"
					}
				default:
					posctl++
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

				}
			}
			switch posctl {
			case 1:
				pos1 = pos
				fmt.Println("test1")

			case 2:
				pos2 = pos
				fmt.Println("test2")
				fmt.Println(pos)
			case 3:
				pos3 = pos
				fmt.Println("test3")

			case 4:
				pos4 = pos

			}
			fmt.Printf("pos1 %s\n", pos1)
			fmt.Printf("pos2 %s\n", pos2)
			fmt.Printf("pos3 %s\n", pos3)
			fmt.Printf("pos4 %s\n", pos4)

		case term.EventError:
			panic(ev.Err)
		}
		switch level {
		case 1:
			if tracktype == 1 {
				fmt.Println("Mapping Straight-Line Track")
				fmt.Println("Select First GPS Point")
				fmt.Println("(N)orth (S)outh (E)ast (W)est")
				fmt.Println("Enter Selection:")
				level = 2
			}
			if tracktype == 2 {
				fmt.Println("Mapping Square Four Point Track")
			}

		case 2:
			if tracktype == 1 {
				fmt.Println("Mapping Straight-Line Track")
				fmt.Println("Select Second GPS Point")
				fmt.Println("(N)orth (S)outh (E)ast (W)est")
				fmt.Println("Enter Selection:")
				level = 3
			}
			if tracktype == 2 {
				fmt.Println("Mapping Square Four Point Track")
			}

		case 3:
			if tracktype == 1 {
				fmt.Println("Mapping Straight-Line Track")
				fmt.Println("Save Track Map")
				fmt.Println("(Y)es (N)o")
				level = 4
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
