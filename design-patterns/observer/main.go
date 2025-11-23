package main

import "fmt"

type Observer interface {
	Update(string)
}

type Product struct {
	price     string
	observers []Observer
}

func (p *Product) RegisterObserver(o Observer) {
	p.observers = append(p.observers, o)
}

func (p *Product) setPrice(price string) {
	p.price = price
	p.NotifyObservers()
}

func (p *Product) NotifyObservers() {
	for _, observer := range p.observers {
		observer.Update(p.price)
	}
}

type PriceDisplay struct {
	Name string
}

func (p *PriceDisplay) Update(price string) {
	fmt.Printf("%s : Price updated to %s\n", p.Name, price)
}

func main() {

	product := &Product{}

	display1 := &PriceDisplay{Name: "Display 1"}
	display2 := &PriceDisplay{Name: "Display 2"}

	product.RegisterObserver(display1)
	product.RegisterObserver(display2)

	product.setPrice("20.000T")
	product.setPrice("25.000T")

}
