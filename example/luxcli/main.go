package main

import (
	"fmt"
	"os"

	"github.com/leosunmo/goluxafor"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("provide colour as arg. Colours: red, green, blue, reset|off|black")
		os.Exit(1)
	}
	colour := os.Args[1]
	luxafor := goluxafor.NewLuxafor()

	switch colour {
	case "red":
		luxafor.Colour(goluxafor.LedAll, 255, 0, 0, 0)
	case "green":
		luxafor.Colour(goluxafor.LedAll, 0, 255, 0, 0)
	case "blue":
		luxafor.Colour(goluxafor.LedAll, 0, 0, 255, 0)
	case "reset", "black", "off":
		luxafor.Colour(goluxafor.LedAll, 0, 0, 0, 0)
	default:
		fmt.Printf("%s is not a supported colour", colour)
	}

	luxafor.Close()
}
