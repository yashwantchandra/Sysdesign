package main

import "fmt"

type Observer interface {
	Update(msg string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type UserObserver struct {
	name string
}

func (u UserObserver) Update(msg string) {
	fmt.Printf("[%s] received: %s\n", u.name, msg)
}

func main() {
	sub := &Subject{}
	sub.Attach(UserObserver{name: "Alice"})
	sub.Attach(UserObserver{name: "Bob"})

	sub.Notify("Sale started!")
}
