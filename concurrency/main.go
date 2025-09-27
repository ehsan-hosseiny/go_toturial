package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go f1()
	go f2()
	go f3()

	wg.Add(3)
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

}

func f1() {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Println("for-f1", i)
	}
}

func f2() {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Println("for-f2", i)
	}
}

func f3() {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Println("for-f3", i)
	}
}
