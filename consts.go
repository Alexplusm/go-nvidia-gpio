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
