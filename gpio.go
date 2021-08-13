package gonvdgpio

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

type Pin struct {
	indexNumber int
	sysfsNumber int
}

func SetupPin(indexNumber int, direction string, level int) (*Pin, error) {
	sysfsNumber, ok := gpioIndexNumberToSysfsNumberMap[indexNumber]
	if !ok {
		return nil, fmt.Errorf("gonvdgpio[.SetupPin]: invalid pin number: %v", indexNumber)
	}

	pin := &Pin{indexNumber: indexNumber, sysfsNumber: sysfsNumber}

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

func (p *Pin) SetLevel(level int) error {
	levelDirectory := path.Join(gpiosDir, p.getDirectoryName(), "value")

	err := ioutil.WriteFile(levelDirectory, []byte(strconv.Itoa(level)), 0666)
	if err != nil {
		return fmt.Errorf("[Pin.SetLevel]: %+v", err)
	}

	return nil
}

func (p *Pin) GetLevel() (int, error) {
	levelDir := path.Join(gpiosDir, p.getDirectoryName(), "value")

	content, err := ioutil.ReadFile(levelDir)
	if err != nil {
		return 0, fmt.Errorf("[Pin.GetLevel][1]: %+v", err)
	}

	levelRaw := strings.TrimRight(string(content), "\n")

	fmt.Println("levelRaw", levelRaw, "|", string(content))

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

func (p *Pin) SetDirection(direction string) error {
	directionPath := path.Join(gpiosDir, p.getDirectoryName(), "direction")

	err := ioutil.WriteFile(directionPath, []byte(direction), 0666)
	if err != nil {
		return fmt.Errorf("[Pin.SetDirection]: %+v", err)
	}

	return nil
}

func (p *Pin) GetDirection() (string, error) {
	directionPath := path.Join(gpiosDir, p.getDirectoryName(), "direction")

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

func ListPins() {
	panic("implements")
}

func CleanUp() error {
	panic("implements")
}

// ---

func (p *Pin) setup() error {
	filePath := path.Join(gpiosDir, "export")

	err := ioutil.WriteFile(filePath, []byte(strconv.Itoa(p.sysfsNumber)), 0666)
	if err != nil {
		return fmt.Errorf("[Pin.setup]: %+v", err)
	}

	return nil
}

func (p Pin) getDirectoryName() string {
	return fmt.Sprintf("gpio%v", p.sysfsNumber)
}
