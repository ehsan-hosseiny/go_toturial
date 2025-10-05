package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	numImages := 10

	downloadChannel := make(chan string, numImages)

	processedChannel := make(chan string, numImages)

	for i := 0; i < numImages; i++ {
		//Fan out step
		wg.Add(1)
		go downloadImage(i, downloadChannel)

	}

	go func() {
		wg.Wait()
		close(downloadChannel)
	}()

	go func() {
		processImage(downloadChannel, processedChannel)
		close(processedChannel)
	}()

	for res := range processedChannel {
		fmt.Println(res)
	}

	fmt.Println("All images proccessed.")

}

func downloadImage(imageID int, out chan<- string) {
	defer wg.Done()
	time.Sleep(time.Microsecond * 500)
	out <- fmt.Sprintf("Image %d downloaded ", imageID)
}

func processImage(imageRes <-chan string, out chan<- string) {
	for res := range imageRes {
		out <- fmt.Sprintf("%s proccessed.", res)

	}

}
