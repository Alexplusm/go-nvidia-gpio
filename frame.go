package gonvdgpio

// TODO: implements if need

//type FrameConfigItem struct {
//	Pin       int
//	Direction string
//	Level     int
//}
//
//type FrameConfig = []FrameConfigItem
//
//type PinFrame struct {
//	pins []Pin
//}
//
//func BuildFrame(config FrameConfig) (*PinFrame, error) {
//	frame := new(PinFrame)
//	for _, item := range config {
//		pin, err := NewPin(item.Pin, item.Direction, item.Level)
//		if err != nil {
//			return nil, fmt.Errorf("[BuildFrame]: %v", err)
//		}
//
//		frame.pins = append(frame.pins, *pin)
//	}
//
//	return frame, nil
//}
//
//func (f *PinFrame) SetLevels(levelConfig string) error {
//	if len(levelConfig) != len(f.pins) {
//		return fmt.Errorf("[PinFrame.SetLevels][1]: invalid level config")
//	}
//
//	for index, level := range levelConfig {
//		lvl, err := strconv.Atoi(string(level))
//		if err != nil {
//			return fmt.Errorf("[PinFrame.SetLevels][2]: %+v", err)
//		}
//
//		err = f.pins[index].SetLevel(lvl)
//		if err != nil {
//			return fmt.Errorf("[PinFrame.SetLevels][3]: %+v", err)
//		}
//	}
//
//	return nil
//}
