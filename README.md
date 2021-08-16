##### Climb into the root
akshay@jetson-nano:~$ sudo su
##### Set the GPIO
##### gpio79 is pin 12 on the Jetson Nano Sysfs
root@jetson-nano:~$ echo 79 > /sys/class/gpio/export
##### Set if the pin is an Output or Input pin
root@jetson-nano:~$ echo out > /sys/class/gpio/gpio79/direction
##### Set the digital pin 1 (HIGH) or 0 (LOW)
root@jetson-nano:~$ echo 1 > /sys/class/gpio/gpio79/value
root@jetson-nano:~$ echo 0 > /sys/class/gpio/gpio79/value

https://maker.pro/nvidia-jetson/tutorial/how-to-use-gpio-pins-on-jetson-nano-developer-kit

// INFO: SRC: https://gist.github.com/titpetric/bc3e43b2c6efc2cd9364cf52ffbc17c4

# Release

* git tag v1.0.0
* git push origin v1.0.0
