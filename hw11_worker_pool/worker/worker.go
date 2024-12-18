package worker

import (
	"fmt"
	"sync"
)

func DoWork(i int, counter *int, mtx *sync.Mutex, wg *sync.WaitGroup) {
	defer mtx.Unlock()
	defer wg.Done()
	mtx.Lock()
	c := *counter
	c++
	*counter = c
	fmt.Printf("Worker %d: %d\n", i, *counter)
}
