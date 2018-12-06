package main

import (
	"math/rand"
	"time"

	"github.com/leosunmo/goluxafor"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	luxafor := goluxafor.NewLuxafor()
	// luxafor.Strobe(goluxafor.LedAll, goluxafor.RandomColour(), goluxafor.RandomColour(), goluxafor.RandomColour(), 10, 2)
	luxafor.Wave(goluxafor.Wave1, 255, 0, 255, 1, 10)
	// luxafor.Pattern(goluxafor.Pattern4, 0)
	// luxafor.Colour(goluxafor.Led1, goluxafor.RandomColour(), goluxafor.RandomColour(), goluxafor.RandomColour(), 0)
	// luxafor.Colour(goluxafor.Led2, goluxafor.RandomColour(), goluxafor.RandomColour(), goluxafor.RandomColour(), 0)
	// luxafor.Colour(goluxafor.Led3, goluxafor.RandomColour(), goluxafor.RandomColour(), goluxafor.RandomColour(), 0)
	// luxafor.Colour(goluxafor.Led4, goluxafor.RandomColour(), goluxafor.RandomColour(), goluxafor.RandomColour(), 0)
	// luxafor.Colour(goluxafor.Led5, goluxafor.RandomColour(), goluxafor.RandomColour(), goluxafor.RandomColour(), 0)
	// luxafor.Colour(goluxafor.Led6, goluxafor.RandomColour(), goluxafor.RandomColour(), goluxafor.RandomColour(), 0)
	// luxafor.Colour(goluxafor.Led1, 0, 200, 0, 0)
	// luxafor.Colour(goluxafor.Led2, 0, 0, 0, 0)
	// luxafor.Colour(goluxafor.Led3, 255, 0, 0, 0)
	// luxafor.Colour(goluxafor.Led4, 0, 200, 0, 0)
	// luxafor.Colour(goluxafor.Led5, 0, 0, 0, 0)
	// luxafor.Colour(goluxafor.Led6, 255, 0, 0, 0)
	luxafor.Close()
}
