package main

import (
	"fmt"
	"sync"

	"github.com/Bladforceone/go_hw_otus/hw11_worker_pool/worker"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	counter := 0
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			defer mutex.Unlock()
			mutex.Lock()
			worker.DoWork(&counter)
			fmt.Printf("Worker %d: %d\n", i, counter)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Counter: %d\n", counter)
}
