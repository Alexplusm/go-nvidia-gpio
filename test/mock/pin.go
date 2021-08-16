package mock

import (
	"fmt"

	gonvdgpio "github.com/alexplusm/go-nvidia-gpio"
)

type pin struct {
	level     int
	direction string
}

//	TODO: strategy for "IN" direction (chan ?)
func NewPin(indexNumber int, direction string, level int) (gonvdgpio.IPin, error) {
	if ok := gonvdgpio.IndexPinNumberExist(indexNumber); !ok {
		return nil, fmt.Errorf("gonvdgpio[.NewPin]: pin index number invalid: %v", indexNumber)
	}

	return &pin{level: level, direction: direction}, nil
}

func (p pin) SetLevel(level int) (err error) {
	switch level {
	case gonvdgpio.HIGH:
		return
	case gonvdgpio.LOW:
		return
	}

	return fmt.Errorf("[pin.SetLevel]: undefined")
}

func (p pin) GetLevel() (int, error) {
	return p.level, nil
}

func (p pin) SetDirection(direction string) (err error) {
	switch direction {
	case gonvdgpio.IN:
		return
	case gonvdgpio.OUT:
		return
	}

	return fmt.Errorf("[pin.SetDirection]: undefined")
}

func (p pin) GetDirection() (string, error) {
	return p.direction, nil
}

func (p pin) Unexport() (err error) {
	return
}
