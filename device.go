package goluxafor

import (
	"log"

	"github.com/karalabe/hid"
)

const vendorId = uint16(0x4d8)   // Microchip Technology Inc.
const productId = uint16(0xf372) // LUXAFOR FLAG

type Led byte

const (
	LedAll Led = 0xff
	LedA   Led = 0x41
	LedB   Led = 0x42
	Led1   Led = 0x01
	Led2   Led = 0x02
	Led3   Led = 0x03
	Led4   Led = 0x04
	Led5   Led = 0x05
	Led6   Led = 0x06
)

type Wave byte

const (
	Wave1 Wave = 0x01
	Wave2 Wave = 0x02
	Wave3 Wave = 0x03
	Wave4 Wave = 0x04
	Wave5 Wave = 0x05
)

type Pattern byte

const (
	Pattern1 Pattern = 0x01
	Pattern2 Pattern = 0x02
	Pattern3 Pattern = 0x03
	Pattern4 Pattern = 0x04
	Pattern5 Pattern = 0x05
	Pattern6 Pattern = 0x06
	Pattern7 Pattern = 0x07
	Pattern8 Pattern = 0x08
)

type ProdColour byte

const (
	Enable  ProdColour = 0x45
	Disable ProdColour = 0x44
	Red     ProdColour = 0x52
	Green   ProdColour = 0x47
	Blue    ProdColour = 0x42
	Cyan    ProdColour = 0x43
	Magenta ProdColour = 0x4d
	Yellow  ProdColour = 0x59
	White   ProdColour = 0x57
	Off     ProdColour = 0x4f
)

func newDevice() *hid.Device {

	devInfo := hid.Enumerate(vendorId, productId)

	if len(devInfo) < 1 {
		log.Fatalf("no devices found matching VID %v and PID %v", vendorId, productId)
	}
	if len(devInfo) > 1 {
		log.Fatalf("More than one device found matching VID %v and PID %v", vendorId, productId)
	}

	dev, err := devInfo[0].Open()
	if err != nil {
		log.Fatalf("Failed to open HID device, err: %s", err)
	}

	return dev
}
