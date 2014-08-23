package main

import (
	"github.com/edmontongo/gobot"
	"github.com/edmontongo/gobot/platforms/firmata"
	"github.com/edmontongo/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "/dev/ttyACM0")
	button := gpio.NewMakeyButtonDriver(firmataAdaptor, "button", "2")
	led := gpio.NewLedDriver(firmataAdaptor, "led", "13")

	work := func() {
		gobot.On(button.Event("push"), func(data interface{}) {
			led.On()
		})

		gobot.On(button.Event("release"), func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("makeyBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{button, led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
