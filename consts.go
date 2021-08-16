package gonvdgpio

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

	sysfsValueDir     = "value"
	sysfsDirectionDir = "direction"
	sysfsExportDir    = "export"
	sysfsUnexportDir  = "unexport"
)

var gpioIndexNumberToSysfsNumberMap map[int]int

func init() {
	gpioIndexNumberToSysfsNumberMap = map[int]int{
		7:  216,
		11: 50,
		12: 79,
		13: 14,
		15: 194,
		16: 232,
		18: 15,
		19: 16,
		21: 17,
		22: 13,
		23: 18,
		24: 19,
		26: 20,
		29: 149,
		31: 200,
		32: 168,
		33: 38,
		35: 76,
		36: 51,
		37: 12,
		38: 77,
		40: 78,
	}
}
