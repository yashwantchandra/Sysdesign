package main

import "fmt"

type Coffee interface {
	Cost() int
	Description() string
}

type SimpleCoffee struct{}

func (c SimpleCoffee) Cost() int              { return 10 }
func (c SimpleCoffee) Description() string   { return "Simple coffee" }

type MilkDecorator struct {
	coffee Coffee
}

func (m MilkDecorator) Cost() int {
	return m.coffee.Cost() + 2
}

func (m MilkDecorator) Description() string {
	return m.coffee.Description() + ", milk"
}

type SugarDecorator struct {
	coffee Coffee
}

func (s SugarDecorator) Cost() int {
	return s.coffee.Cost() + 1
}

func (s SugarDecorator) Description() string {
	return s.coffee.Description() + ", sugar"
}

func main() {
	var c Coffee = SimpleCoffee{}
	c = MilkDecorator{coffee: c}
	c = SugarDecorator{coffee: c}

	fmt.Println(c.Description(), "->", c.Cost(), "cents")
}
