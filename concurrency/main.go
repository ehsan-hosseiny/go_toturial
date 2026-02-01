package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go f1()
	go f2()
	go f3()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

}

func f1() {
	defer wg.Done()
	
	for i := 0; i < 5; i++ {
		fmt.Println("for-a", i)
	}
}

func f2() {
	defer wg.Done()
	
	for i := 0; i < 5; i++ {
		fmt.Println("for-b", i)
	}

}

func f3() {
	defer wg.Done()
	
	for i := 0; i < 5; i++ {
		fmt.Println("for-c", i)
	}
}
