package gonvdgpio

const (
	gpioSysfsNumber216 = 216
	gpioSysfsNumber50  = 50
	gpioSysfsNumber14  = 14
	gpioSysfsNumber194 = 194
	gpioSysfsNumber16  = 16
	gpioSysfsNumber17  = 17
	gpioSysfsNumber18  = 18
	gpioSysfsNumber149 = 149
	gpioSysfsNumber200 = 200
	gpioSysfsNumber38  = 38
	gpioSysfsNumber76  = 76
	gpioSysfsNumber12  = 12

	gpioSysfsNumber79  = 79
	gpioSysfsNumber232 = 232
	gpioSysfsNumber15  = 15
	gpioSysfsNumber13  = 13
	gpioSysfsNumber19  = 19
	gpioSysfsNumber20  = 20
	gpioSysfsNumber168 = 168
	gpioSysfsNumber51  = 51
	gpioSysfsNumber77  = 77
	gpioSysfsNumber78  = 78
)

const (
	LOW = iota
	HIGH
)

const (
	IN  = "in"
	OUT = "out"
)

const (
	gpiosDir = "/sys/class/gpio"
)

const (
	sysfsValue     = "value"
	sysfsDirection = "direction"
	sysfsExport    = "export"
	sysfsUnexport  = "unexport"
)

var gpioIndexNumberToSysfsNumberMap map[int]int

func init() {
	gpioIndexNumberToSysfsNumberMap = map[int]int{
		7:  gpioSysfsNumber216,
		11: gpioSysfsNumber50,
		12: gpioSysfsNumber79,
		13: gpioSysfsNumber14,
		15: gpioSysfsNumber194,
		16: gpioSysfsNumber232,
		18: gpioSysfsNumber15,
		19: gpioSysfsNumber16,
		21: gpioSysfsNumber17,
		22: gpioSysfsNumber13,
		23: gpioSysfsNumber18,
		24: gpioSysfsNumber19,
		26: gpioSysfsNumber20,
		29: gpioSysfsNumber149,
		31: gpioSysfsNumber200,
		32: gpioSysfsNumber168,
		33: gpioSysfsNumber38,
		35: gpioSysfsNumber76,
		36: gpioSysfsNumber51,
		37: gpioSysfsNumber12,
		38: gpioSysfsNumber77,
		40: gpioSysfsNumber78,
	}
}
