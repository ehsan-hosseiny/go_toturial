package main

import "sync"

var wg sync.WaitGroup

func main() {

	ch := make(chan int, 5)

	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()

}

func f1(ch chan int) {
	defer wg.Done()
	ch <- 10
	ch <- 20 

}

func f2(ch chan int) {
	defer wg.Done()

}
