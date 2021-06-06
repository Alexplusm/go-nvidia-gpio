package gonvdgpio

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

type Pin struct {
	Value int
}

func (p *Pin) setValue(value int) *Pin {
	p.Value = value

	return p
}

func (p *Pin) setup() error {
	filePath := path.Join(gpiosDirectory, "export")

	err := ioutil.WriteFile(filePath, []byte(strconv.Itoa(p.Value)), 0666)
	if err != nil {
		return fmt.Errorf("[Pin.setup]: %+v", err)
	}

	return nil
}

func (p *Pin) getDirectoryName() string {
	return fmt.Sprintf("gpio%v", p.Value)
}

func (p *Pin) SetDirection(direction string) error {
	directionPath := path.Join(gpiosDirectory, p.getDirectoryName(), "direction")

	err := ioutil.WriteFile(directionPath, []byte(direction), 0666)
	if err != nil {
		return fmt.Errorf("[Pin.SetDirection]: %+v", err)
	}

	return nil
}

func (p *Pin) GetDirection() (string, error) {
	directionPath := path.Join(gpiosDirectory, p.getDirectoryName(), "direction")

	content, err := ioutil.ReadFile(directionPath)
	if err != nil {
		return "", fmt.Errorf("[Pin.GetDirection]: %v", err)
	}

	switch string(content) {
	case IN:
		return IN, nil
	case OUT:
		return OUT, nil
	}

	return "", fmt.Errorf("[Pin.GetDirection]: undefined")
}

func (p *Pin) SetLevel(level int) error {
	levelDirectory := path.Join(gpiosDirectory, p.getDirectoryName(), "value")

	err := ioutil.WriteFile(levelDirectory, []byte(strconv.Itoa(level)), 0666)
	if err != nil {
		return fmt.Errorf("[Pin.SetLevel]: %+v", err)
	}

	return nil
}

func (p *Pin) GetLevel() (int, error) {
	levelDirectory := path.Join(gpiosDirectory, p.getDirectoryName(), "value")

	content, err := ioutil.ReadFile(levelDirectory)
	if err != nil {
		return 0, fmt.Errorf("[Pin.GetLevel][1]: %+v", err)
	}

	levelRaw := strings.TrimRight(string(content), "\n")

	level, err := strconv.Atoi(levelRaw)
	if err != nil {
		return 0, fmt.Errorf("[Pin.GetLevel][2]: %+v", err)
	}

	switch level {
	case LOW:
		return LOW, nil
	case HIGH:
		return HIGH, nil
	}

	return 0, fmt.Errorf("[Pin.GetLevel][3]: undefined")
}

// TODO
func ListPins() {}

// TODO
func CleanUp() error {
	return nil
}

func SetupPin(pinNumber int, direction string, level int) (*Pin, error) {
	pin := new(Pin).setValue(pinNumber)

	if err := pin.setup(); err != nil {
		// TODO
	}

	if err := pin.SetDirection(direction); err != nil {
		// TODO
	}

	if err := pin.SetLevel(level); err != nil {
		// TODO
	}

	return pin, nil
}
