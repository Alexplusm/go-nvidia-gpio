package gonvdgpio

import (
	"bytes"
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
		return nil, fmt.Errorf("gonvdgpio[.SetupPin][1]: invalid pin number: %v", indexNumber)
	}

	pin := &Pin{indexNumber: indexNumber, sysfsNumber: sysfsNumber}

	if err := pin.setup(); err != nil {
		return nil, fmt.Errorf("gonvdgpio[.SetupPin][2]: %v", err)
	}

	if err := pin.SetDirection(direction); err != nil {
		return nil, fmt.Errorf("gonvdgpio[.SetupPin][3]: %v", err)
	}

	if err := pin.SetLevel(level); err != nil {
		return nil, fmt.Errorf("gonvdgpio[.SetupPin][4]: %v", err)
	}

	return pin, nil
}

func (p Pin) SetLevel(level int) (err error) {
	dir := path.Join(gpiosDir, p.getSysfsGpioPinName(), sysfsValue)

	if err = ioutil.WriteFile(dir, []byte(strconv.Itoa(level)), 0666); err != nil {
		return fmt.Errorf("gonvdgpio[Pin.SetLevel][1]: %+v", err)
	}

	return nil
}

func (p Pin) GetLevel() (int, error) {
	dir := path.Join(gpiosDir, p.getSysfsGpioPinName(), sysfsValue)

	content, err := ioutil.ReadFile(dir)
	if err != nil {
		return 0, fmt.Errorf("gonvdgpio[Pin.GetLevel][1]: %+v", err)
	}

	levelRaw := strings.TrimRight(string(content), "\n")

	level, err := strconv.Atoi(levelRaw)
	if err != nil {
		return 0, fmt.Errorf("gonvdgpio[Pin.GetLevel][2]: %+v", err)
	}

	switch level {
	case LOW:
		return LOW, nil
	case HIGH:
		return HIGH, nil
	}

	return 0, fmt.Errorf("gonvdgpio[Pin.GetLevel][3]: undefined")
}

func (p Pin) SetDirection(direction string) error {
	dir := path.Join(gpiosDir, p.getSysfsGpioPinName(), sysfsDirection)

	err := ioutil.WriteFile(dir, []byte(direction), 0666)
	if err != nil {
		return fmt.Errorf("gonvdgpio[Pin.SetDirection][1]: %+v", err)
	}

	return nil
}

func (p Pin) GetDirection() (string, error) {
	dir := path.Join(gpiosDir, p.getSysfsGpioPinName(), sysfsDirection)

	content, err := ioutil.ReadFile(dir)
	if err != nil {
		return "", fmt.Errorf("gonvdgpio[Pin.GetDirection][1]: %v | %v", err, err.Error())
	}

	content = bytes.TrimSpace(content)

	switch string(content) {
	case IN:
		return IN, nil
	case OUT:
		return OUT, nil
	}

	return "", fmt.Errorf("gonvdgpio[Pin.GetDirection][1]: undefined")
}

func (p *Pin) Unexport() (err error) {
	dir := path.Join(gpiosDir, sysfsUnexport)

	if err = ioutil.WriteFile(dir, []byte(strconv.Itoa(p.sysfsNumber)), 0666); err != nil {
		return fmt.Errorf("gonvdgpio[Pin.Unexport][1]: %+v", err.Error())
	}

	p.indexNumber = -1
	p.sysfsNumber = -1

	return
}

// --- private

func (p Pin) setup() (err error) {
	filePath := path.Join(gpiosDir, sysfsExport)

	if err = ioutil.WriteFile(filePath, []byte(strconv.Itoa(p.sysfsNumber)), 0666); err != nil {
		return fmt.Errorf("[Pin.setup][1]: %+v", err)
	}

	return nil
}

func (p Pin) getSysfsGpioPinName() string {
	return fmt.Sprintf("gpio%v", p.sysfsNumber)
}
