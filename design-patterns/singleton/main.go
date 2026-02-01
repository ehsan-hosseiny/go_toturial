package main

<<<<<<< HEAD
import (
	"fmt"
	"sync"
)

type Singleton struct {
	Value string
}

var instance *Singleton
=======
import "sync"

type Singleton struct {
	value string
}

var instance *Singleton

>>>>>>> 6815f6c (added design pattern)
var once sync.Once

func GetInstance(value string) *Singleton {
	once.Do(func() {
<<<<<<< HEAD
		instance = &Singleton{
			Value: value,
		}
	})

	return instance

}

func main() {
	s1 := GetInstance("First Instance")
	fmt.Println(s1)

	s2 := GetInstance("Second Instance")
	fmt.Println(s2)
=======
		instance = &Singleton{value: value}
	})
	return instance
}

func main() {
	s1 := GetInstance("first instance")
	s2 := GetInstance("second instance")

	println(s1.value) // Output: first
	println(s2.value) // Output: first
>>>>>>> 6815f6c (added design pattern)

}
