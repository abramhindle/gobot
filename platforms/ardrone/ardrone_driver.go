package ardrone

import (
	"github.com/edmontongo/gobot"
)

type ArdroneDriver struct {
	gobot.Driver
}

func NewArdroneDriver(adaptor *ArdroneAdaptor, name string) *ArdroneDriver {
	d := &ArdroneDriver{
		Driver: *gobot.NewDriver(
			name,
			"ArdroneDriver",
			adaptor,
		),
	}
	d.AddEvent("flying")
	return d
}
func (a *ArdroneDriver) adaptor() *ArdroneAdaptor {
	return a.Adaptor().(*ArdroneAdaptor)
}

func (a *ArdroneDriver) Start() bool {
	return true
}

func (a *ArdroneDriver) Halt() bool {
	return true
}

func (a *ArdroneDriver) TakeOff() {
	gobot.Publish(a.Event("flying"), a.adaptor().drone.Takeoff())
}

func (a *ArdroneDriver) Land() {
	a.adaptor().drone.Land()
}

func (a *ArdroneDriver) Up(n float64) {
	a.adaptor().drone.Up(n)
}

func (a *ArdroneDriver) Down(n float64) {
	a.adaptor().drone.Down(n)
}

func (a *ArdroneDriver) Left(n float64) {
	a.adaptor().drone.Left(n)
}

func (a *ArdroneDriver) Right(n float64) {
	a.adaptor().drone.Right(n)
}

func (a *ArdroneDriver) Forward(n float64) {
	a.adaptor().drone.Forward(n)
}

func (a *ArdroneDriver) Backward(n float64) {
	a.adaptor().drone.Backward(n)
}

func (a *ArdroneDriver) Clockwise(n float64) {
	a.adaptor().drone.Clockwise(n)
}

func (a *ArdroneDriver) CounterClockwise(n float64) {
	a.adaptor().drone.Counterclockwise(n)
}

func (a *ArdroneDriver) Hover() {
	a.adaptor().drone.Hover()
}
