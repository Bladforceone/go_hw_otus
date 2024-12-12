package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw10_motion_sensor/sensor"
)

func main() {
	sens := make(chan int)
	proc := make(chan int)
	go sensor.SensorData(sens)
	go sensor.ProcessData(sens, proc)
	i := 0
	for data := range proc {
		i++
		fmt.Println(data, " ", i)
	}
}
