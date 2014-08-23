package gpio

import (
	"github.com/edmontongo/gobot"
)

type LedDriver struct {
	gobot.Driver
	High bool
}

func NewLedDriver(a PwmDigitalWriter, name string, pin string) *LedDriver {

	l := &LedDriver{
		Driver: *gobot.NewDriver(
			name,
			"LedDriver",
			pin,
			a.(gobot.AdaptorInterface),
		),
		High: false,
	}

	l.AddCommand("Brightness", func(params map[string]interface{}) interface{} {
		level := byte(params["level"].(float64))
		l.Brightness(level)
		return nil
	})

	l.AddCommand("Toggle", func(params map[string]interface{}) interface{} {
		l.Toggle()
		return nil
	})

	l.AddCommand("On", func(params map[string]interface{}) interface{} {
		l.On()
		return nil
	})

	l.AddCommand("Off", func(params map[string]interface{}) interface{} {
		l.Off()
		return nil
	})

	return l
}

func (l *LedDriver) adaptor() PwmDigitalWriter {
	return l.Adaptor().(PwmDigitalWriter)
}

func (l *LedDriver) Start() bool { return true }
func (l *LedDriver) Halt() bool  { return true }
func (l *LedDriver) Init() bool  { return true }

func (l *LedDriver) IsOn() bool {
	return l.High
}

func (l *LedDriver) IsOff() bool {
	return !l.IsOn()
}

func (l *LedDriver) On() bool {
	l.changeState(1)
	l.High = true
	return true
}

func (l *LedDriver) Off() bool {
	l.changeState(0)
	l.High = false
	return true
}

func (l *LedDriver) Toggle() {
	if l.IsOn() {
		l.Off()
	} else {
		l.On()
	}
}

func (l *LedDriver) Brightness(level byte) {
	l.adaptor().PwmWrite(l.Pin(), level)
}

func (l *LedDriver) changeState(level byte) {
	l.adaptor().DigitalWrite(l.Pin(), level)
}
