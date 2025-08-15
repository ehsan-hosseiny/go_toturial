package main

import "fmt"

func main() {
	slice := []int{10, 12, 44, 63, 76, 101, 11, 3, 72, 54}

	// res := getElement(slice, 2)
	// fmt.Println(res)

	// fmt.Println(findMax(slice))
	bubbleSort(slice)
	fmt.Println(slice)

}

// o(1) example
func getElement(slice []int, index int) int {
	return slice[index]

}

// o(n) example
func findMax(slice []int) int {
	max := slice[0]

	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max

}

// o(n^2) example
func bubbleSort(slice []int) {
	n := len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

}
