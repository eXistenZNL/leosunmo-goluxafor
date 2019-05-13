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
	pattern = flag.Int("p", 0, "pattern to activate (1-8)")

	rgb     = flag.String("rgb", "", "`RGB` value to use (in web #C0C0C0 format)")
	fade    = flag.Uint("fade", 0, "`fade` value")
	ledmask = flag.Uint("leds", uint(goluxafor.LedAll), "`led mask value` to use to apply the colour")

	off = flag.Bool("off", false, "switch off")
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
		err = luxafor.Pattern(goluxafor.Pattern(*pattern), 0)
	case *rgb != "":
		c, err := ParseHexColor(*rgb)
		if err != nil {
			break
		}
		err = luxafor.Colour(goluxafor.Led(*ledmask), c.R, c.G, c.B, byte(*fade))
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
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}
	return
}
