package gonvdgpio

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
)

type IPin interface {
	SetLevel(level int) (err error)
	GetLevel() (int, error)
	SetDirection(direction string) (err error)
	GetDirection() (string, error)
	Unexport() (err error)
}

type pin struct {
	indexNumber int
	sysfsNumber int
}

func NewPin(indexNumber int, direction string, level int) (IPin, error) {
	sysfsNumber, ok := gpioIndexNumberToSysfsNumberMap[indexNumber]
	if !ok {
		return nil, fmt.Errorf("gonvdgpio[.NewPin][1]: invalid pin number: %v", indexNumber)
	}

	p := &pin{indexNumber: indexNumber, sysfsNumber: sysfsNumber}

	if err := p.setup(); err != nil {
		return nil, fmt.Errorf("gonvdgpio[.NewPin][2]: %v", err)
	}

	if err := p.SetDirection(direction); err != nil {
		return nil, fmt.Errorf("gonvdgpio[.NewPin][3]: %v", err)
	}

	if err := p.SetLevel(level); err != nil {
		return nil, fmt.Errorf("gonvdgpio[.NewPin][4]: %v", err)
	}

	return p, nil
}

func (p pin) SetLevel(level int) (err error) {
	dir := path.Join(gpiosDir, p.getSysfsGpioPinName(), sysfsValueDir)

	if err = ioutil.WriteFile(dir, []byte(strconv.Itoa(level)), 0666); err != nil {
		return fmt.Errorf("gonvdgpio[pin.SetLevel][1]: %+v", err)
	}

	return
}

func (p pin) GetLevel() (int, error) {
	dir := path.Join(gpiosDir, p.getSysfsGpioPinName(), sysfsValueDir)

	content, err := ioutil.ReadFile(dir)
	if err != nil {
		return 0, fmt.Errorf("gonvdgpio[pin.GetLevel][1]: %+v", err)
	}

	content = bytes.TrimSpace(content)

	level, err := strconv.Atoi(string(content))
	if err != nil {
		return 0, fmt.Errorf("gonvdgpio[pin.GetLevel][2]: %+v", err)
	}

	switch level {
	case LOW:
		return LOW, nil
	case HIGH:
		return HIGH, nil
	}

	return 0, fmt.Errorf("gonvdgpio[pin.GetLevel][3]: undefined")
}

func (p pin) SetDirection(direction string) (err error) {
	dir := path.Join(gpiosDir, p.getSysfsGpioPinName(), sysfsDirectionDir)

	if err = ioutil.WriteFile(dir, []byte(direction), 0666); err != nil {
		return fmt.Errorf("gonvdgpio[pin.SetDirection][1]: %+v", err)
	}

	return
}

func (p pin) GetDirection() (string, error) {
	dir := path.Join(gpiosDir, p.getSysfsGpioPinName(), sysfsDirectionDir)

	content, err := ioutil.ReadFile(dir)
	if err != nil {
		return "", fmt.Errorf("gonvdgpio[pin.GetDirection][1]: %v | %v", err, err.Error())
	}

	content = bytes.TrimSpace(content)

	switch string(content) {
	case IN:
		return IN, nil
	case OUT:
		return OUT, nil
	}

	return "", fmt.Errorf("gonvdgpio[pin.GetDirection][1]: undefined")
}

func (p *pin) Unexport() (err error) {
	dir := path.Join(gpiosDir, sysfsUnexportDir)

	if err = ioutil.WriteFile(dir, []byte(strconv.Itoa(p.sysfsNumber)), 0666); err != nil {
		return fmt.Errorf("gonvdgpio[pin.Unexport][1]: %+v", err.Error())
	}

	p.indexNumber = -666
	p.sysfsNumber = -666

	return
}

// --- private

// TODO: rename -> export
func (p pin) setup() (err error) {
	filePath := path.Join(gpiosDir, sysfsExportDir)

	if err = ioutil.WriteFile(filePath, []byte(strconv.Itoa(p.sysfsNumber)), 0666); err != nil {
		return fmt.Errorf("[pin.setup][1]: %+v", err)
	}

	return
}

func (p pin) getSysfsGpioPinName() string {
	return fmt.Sprintf("gpio%v", p.sysfsNumber)
}
