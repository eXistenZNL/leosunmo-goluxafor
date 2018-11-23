package goluxafor

import (
	"log"

	"github.com/karalabe/hid"
)

type Luxafor struct {
	Device *hid.Device
}

func NewLuxafor() Luxafor {

	device := newDevice()

	return Luxafor{
		Device: device,
	}
}

func (l *Luxafor) Close() {
	l.Device.Close()
}

func (l *Luxafor) writeCommand(command []byte) error {
	_, err := l.Device.Write(command)
	if err != nil {
		log.Printf("Error writing data: %s", err)
	}
	return err
}

func (l *Luxafor) Color(led Led, red uint8, green uint8, blue uint8, fadeTime uint8) error {
	data := []byte{0x01, byte(led), red, green, blue, fadeTime, 0x0, 0x0}
	if fadeTime > 0 {
		data[1] = 0x02
	}
	return l.writeCommand(data)
}

func (l *Luxafor) Strobe(led Led, red uint8, green uint8, blue uint8, speed uint8, repeat uint8) error {
	data := []byte{0x03, byte(led), red, green, blue, speed, 0x0, repeat}
	return l.writeCommand(data)
}

func (l *Luxafor) Wave(wave Wave, red uint8, green uint8, blue uint8, speed uint8, repeat uint8) error {
	data := []byte{0x04, byte(wave), red, green, blue, 0x0, repeat, speed}
	return l.writeCommand(data)
}

func (l *Luxafor) Pattern(pattern Pattern, repeat uint8) error {
	data := []byte{0x06, byte(pattern), repeat, 0x0, 0x0, 0x0, 0x0, 0x0}
	return l.writeCommand(data)
}
