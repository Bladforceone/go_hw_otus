package sensor

import (
	"math/rand"
	"time"
)

func SensData(dataChan chan<- int) {
	defer close(dataChan)
	timeout := time.NewTimer(time.Minute)
	for {
		select {
		case dataChan <- rand.Intn(1000): //nolint:gosec
		case <-timeout.C:
			return
		}
	}
}

func ProcessData(dataChan <-chan int, processChan chan<- int) {
	defer close(processChan)
	sum := 0
	c := 0
	for data := range dataChan {
		sum += data
		c++
		if c == 10 {
			processChan <- sum
			sum = 0
			c = 0
		}
	}
}
