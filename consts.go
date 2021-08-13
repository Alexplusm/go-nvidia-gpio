package gonvdgpio

const (
	GPIO_216 = 216
	GPIO_50  = 50
	GPIO_14  = 14
	GPIO_194 = 194
	GPIO_16  = 16
	GPIO_17  = 17
	GPIO_18  = 18
	GPIO_149 = 149
	GPIO_200 = 200
	GPIO_38  = 38
	GPIO_76  = 76
	GPIO_12  = 12

	GPIO_79  = 79
	GPIO_232 = 232
	GPIO_15  = 15
	GPIO_13  = 13
	GPIO_19  = 19
	GPIO_20  = 20
	GPIO_168 = 168
	GPIO_51  = 51
	GPIO_77  = 77
	GPIO_78  = 78
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
	gpiosDirectory = "/sys/class/gpio"
)

var gpioIndexNumberToSysfsNumberMap map[int]int

func init() {
	gpioIndexNumberToSysfsNumberMap = map[int]int{
		7:  GPIO_216,
		11: GPIO_50,
		12: GPIO_79,
		13: GPIO_14,
		15: GPIO_194,
		16: GPIO_232,
		18: GPIO_15,
		19: GPIO_16,
		21: GPIO_17,
		22: GPIO_13,
		23: GPIO_18,
		24: GPIO_19,
		26: GPIO_20,
		29: GPIO_149,
		31: GPIO_200,
		32: GPIO_168,
		33: GPIO_38,
		35: GPIO_76,
		36: GPIO_51,
		37: GPIO_12,
		38: GPIO_77,
		40: GPIO_78,
	}
}
