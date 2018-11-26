package main

import (
	"github.com/leosunmo/goluxafor"
)

func main() {
	luxafor := goluxafor.NewLuxafor()
	luxafor.Pattern(goluxafor.Pattern8, 0)

	luxafor.Close()
}
