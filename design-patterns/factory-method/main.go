package main

import "fmt"

type Product interface {
	Use() string
}

type ProductA struct{}

func (p ProductA) Use() string {
	return "Product A in use"
}

type ProductB struct{}

func (p ProductB) Use() string {
	return "Product B in use"
}

func CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return ProductA{}
	case "B":
		return ProductB{}
	default:
		return nil
	}

}

func main() {
	p1 := CreateProduct("A")

	fmt.Println(p1.Use())

	p2 := CreateProduct("B")
	fmt.Println(p2.Use())

}
