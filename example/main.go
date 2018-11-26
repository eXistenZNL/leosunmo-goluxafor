package main

import (
	"math/rand"
	"time"

	"github.com/leosunmo/goluxafor"
)

func randomColor() uint8 {
	return uint8(rand.Intn(0xff))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	luxafor := goluxafor.NewLuxafor()
	luxafor.Strobe(goluxafor.LedAll, randomColor(), randomColor(), randomColor(), 10, 2)
	// luxafor.Pattern(goluxafor.Pattern4, 0)
	// luxafor.Color(goluxafor.Led1, randomColor(), randomColor(), randomColor(), 0)
	// luxafor.Color(goluxafor.Led2, randomColor(), randomColor(), randomColor(), 0)
	// luxafor.Color(goluxafor.Led3, randomColor(), randomColor(), randomColor(), 0)
	// luxafor.Color(goluxafor.Led4, randomColor(), randomColor(), randomColor(), 0)
	// luxafor.Color(goluxafor.Led5, randomColor(), randomColor(), randomColor(), 0)
	// luxafor.Color(goluxafor.Led6, randomColor(), randomColor(), randomColor(), 0)
	// luxafor.Color(goluxafor.Led1, 0, 200, 0, 0)
	// luxafor.Color(goluxafor.Led2, 0, 0, 0, 0)
	// luxafor.Color(goluxafor.Led3, 255, 0, 0, 0)
	// luxafor.Color(goluxafor.Led4, 0, 200, 0, 0)
	// luxafor.Color(goluxafor.Led5, 0, 0, 0, 0)
	// luxafor.Color(goluxafor.Led6, 255, 0, 0, 0)
	luxafor.Close()
}
