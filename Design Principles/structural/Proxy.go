package main

import "fmt"

type Subject interface {
	Request() string
}

type RealSubject struct{}

func (r *RealSubject) Request() string {
	return "real response"
}

// Proxy: controls access to RealSubject (e.g. lazy init, logging).
type Proxy struct {
	real *RealSubject
}

func (p *Proxy) Request() string {
	if p.real == nil {
		p.real = &RealSubject{}
		fmt.Println("Proxy: created real subject")
	}
	return p.real.Request()
}

func main() {
	var sub Subject = &Proxy{}
	fmt.Println(sub.Request())
	fmt.Println(sub.Request())
}
