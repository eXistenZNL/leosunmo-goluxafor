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
	// luxafor.Color(goluxafor.LedAll, randomColor(), randomColor(), randomColor(), 0)
	// luxafor.Strobe(goluxafor.LedAll, randomColor(), randomColor(), randomColor(), 10, 2)
	// luxafor.Wave(goluxafor.Wave4, randomColor(), randomColor(), randomColor(), 10, 2)
	luxafor.Pattern(goluxafor.Pattern3, 1)
	luxafor.Close()
}
