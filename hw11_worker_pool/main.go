package main

import (
	"fmt"
	"sync"

	"github.com/Bladforceone/go_hw_otus/hw11_worker_pool/worker"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker.DoWork(i, &counter, &mutex, &wg)
	}
	wg.Wait()
	fmt.Printf("Counter: %d\n", counter)
}
