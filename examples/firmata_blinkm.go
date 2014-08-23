package main

import (
	"fmt"
	"time"

	"github.com/edmontongo/gobot"
	"github.com/edmontongo/gobot/platforms/firmata"
	"github.com/edmontongo/gobot/platforms/i2c"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "/dev/ttyACM0")
	blinkm := i2c.NewBlinkMDriver(firmataAdaptor, "blinkm")

	work := func() {
		gobot.Every(3*time.Second, func() {
			r := byte(gobot.Rand(255))
			g := byte(gobot.Rand(255))
			b := byte(gobot.Rand(255))
			blinkm.Rgb(r, g, b)
			fmt.Println("color", blinkm.Color())
		})
	}

	robot := gobot.NewRobot("blinkmBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{blinkm},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
