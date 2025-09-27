package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go inc()
	}

	wg.Wait()
	
	fmt.Println("final :", counter)

}

func inc() {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
	fmt.Println(counter)
}
