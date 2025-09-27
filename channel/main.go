package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int)

	wg.Add(2)
	go f1(ch)
	go f2(ch)
	wg.Wait()
}

func f1(ch chan int) {
	defer wg.Done()

	x := 12
	ch <- x
}

func f2(ch chan int) {
	defer wg.Done()

	fmt.Println(<-ch)

}
