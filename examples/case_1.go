package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	gonvdgpio "github.com/alexplusm/go-nvidia-gpio"
)

func main() {
	var (
		pinNumber int
		direction string
		level     int
	)

	flag.IntVar(&pinNumber, "pn", -1, "pin number")
	flag.StringVar(&direction, "d", gonvdgpio.OUT, "direction")
	flag.IntVar(&level, "l", -1, "level")
	flag.Parse()

	p, err := gonvdgpio.SetupPin(pinNumber, direction, level)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = p.Unexport()
		if err != nil {
			fmt.Println("defer: err: ", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println("sig", sig.String())

			err = p.Unexport()
			if err != nil {
				fmt.Println("interrupt: err: ", err)
			}

			close(c)
			return
		}
	}()

	time.Sleep(time.Second * 2)

	for {
		if err = p.SetLevel(gonvdgpio.HIGH); err != nil {
			fmt.Printf("Err: SetLevel: %+v\n", err)
		}

		printCurrState(p)
		time.Sleep(time.Second * 2)

		if err = p.SetLevel(gonvdgpio.LOW); err != nil {
			fmt.Printf("Err: SetLevel: %+v\n", err)
		}

		printCurrState(p)
		time.Sleep(time.Second * 2)
	}
}

func printCurrState(pin *gonvdgpio.Pin) {
	currLevel, err := pin.GetLevel()
	if err != nil {
		fmt.Printf("Err: GetLevel: %+v\n", err)
	}

	currDir, err := pin.GetDirection()
	if err != nil {
		fmt.Printf("Err: GetDirection: %+v\n", err)
	}

	fmt.Printf("Curr state: level = %v, direction = %v\n", currLevel, currDir)
}
