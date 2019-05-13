package main

import (
	"errors"
	"flag"
	"fmt"
	"image/color"
	"log"

	"github.com/leosunmo/goluxafor"
)

var (

	// major modes
	pattern = flag.Int("p", 0, "pattern to activate (1-8)")
	rgb     = flag.String("rgb", "", "`RGB` value to use (in hex #C0C0C0 format)")
	strobe  = flag.Bool("strobe", false, "strobe mode")
	wave    = flag.Int("wave", 0, "wave pattern (`0-5`), 0 - off")

	ledmask = flag.Int("leds", int(goluxafor.LedAll), "led mask value to use to apply the colour. Valid (`1-6,65,66,255`)")
	speed   = flag.Int("speed", 0, "speed for wave and strobe (`0-255`)")

	off = flag.Bool("off", false, "switch off")
)

var (
	// value specific
	validPattern = intValidator(0, 8)
	validWave    = intValidator(0, 5)

	// generic
	validByte = intValidator(0, 255)
)

func main() {
	flag.Parse()

	luxafor := goluxafor.NewLuxafor()
	defer luxafor.Close()

	if *off {
		*rgb = "#000000"
		*pattern = 0
	}

	var err error
	switch {
	default:
		flag.Usage()
		err = errors.New("invalid parameters")
	case *pattern > 0:
		*pattern = validPattern(*pattern)
		err = luxafor.Pattern(goluxafor.Pattern(*pattern), 0)
	case *rgb != "" || *wave > 0 || *strobe:
		*ledmask = validByte(*ledmask)
		*speed = validByte(*speed)
		*wave = validWave(*wave)

		var c color.RGBA
		c, err = ParseHexColor(*rgb)
		if err != nil {
			break
		}
		switch {
		case *rgb != "":
			err = luxafor.Colour(goluxafor.Led(*ledmask), c.R, c.G, c.B, 0)
		case *wave > 0:
			err = luxafor.Wave(goluxafor.Wave(*wave), c.R, c.G, c.B, byte(*speed), 0)
		case *strobe:
			err = luxafor.Strobe(goluxafor.Led(*ledmask), c.R, c.G, c.B, byte(*speed), 0)
		}

	}
	if err != nil {
		log.Fatal(err)
	}
}

// ParseHexColor convert #123456 format colour to color.RGBA.
// An elegant solution from https://stackoverflow.com/a/54200713.
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid color, must be in #000 or #000000 form")
	}
	return
}

func intValidator(min, max int) func(int) int {
	return func(val int) int {
		valid := val
		switch {
		case val < min:
			valid = min
		case max < val:
			valid = max
		}
		return valid
	}
}
