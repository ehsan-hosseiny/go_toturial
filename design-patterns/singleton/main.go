package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Value string
}

var instance *Singleton
var once sync.Once

func GetInstance(value string) *Singleton {
	once.Do(func() {
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

}
