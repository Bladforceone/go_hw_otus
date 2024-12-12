package sensor

import (
	"math/rand"
	"time"
)

func SensData(dataChan chan<- int) {
	defer close(dataChan)
	for i := 0; i < 600; i++ {
		dataChan <- rand.Intn(1000)
		time.Sleep(100 * time.Millisecond)
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
