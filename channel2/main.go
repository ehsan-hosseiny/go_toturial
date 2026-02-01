package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int)

	wg.Add(3)
	go f1(ch)
	go f2(ch)
	go f3(ch)
	wg.Wait()
}

func f1(ch chan int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

}

func f2(ch chan int) {
	defer wg.Done()

	for v := range ch {
		fmt.Println("Received value from channel in Reader 1:", v)
	}

}

func f3(ch chan int) {
	defer wg.Done()

	for v := range ch {
		fmt.Println("Received value from channel in Reader 2:", v)
	}

}
