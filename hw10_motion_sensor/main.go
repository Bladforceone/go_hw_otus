package main

import (
	"fmt"

	"github.com/Bladforceone/go_hw_otus/hw10_motion_sensor/sensor"
)

func main() {
	sens := make(chan int)
	proc := make(chan int)
	go sensor.SensData(sens)
	go sensor.ProcessData(sens, proc)
	i := 0
	for data := range proc {
		i++
		fmt.Println(data, " ", i)
	}
}
