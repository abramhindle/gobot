# Spark

This package provides the Gobot adaptor for the [Spark Core](https://www.spark.io/)

## Installing
```
go get github.com/edmontongo/gobot && go install github.com/edmontongo/gobot/platforms/spark
```

## Example

```go
package main

import (
	"time"

	"github.com/edmontongo/gobot"
	"github.com/edmontongo/gobot/platforms/gpio"
	"github.com/edmontongo/gobot/platforms/spark"
)

func main() {
	gbot := gobot.NewGobot()

	sparkCore := spark.NewSparkCoreAdaptor("spark", "device_id", "access_token")
	led := gpio.NewLedDriver(sparkCore, "led", "D7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{sparkCore},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
```
