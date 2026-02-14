package solid
package main

import "fmt"

// Bird is a more general concept.
type Bird interface {
	Walk()
}

type FlyingBird interface {
	Bird
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Walk() {
	fmt.Println("Sparrow is walking")
}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

type Penguin struct{}

func (p Penguin) Walk() {
	fmt.Println("Penguin is walking")
	// Penguins do NOT fly, so they don't implement Fly.
}

func LetBirdWalk(b Bird) {
	b.Walk()
}

func LetBirdFly(fb FlyingBird) {
	fb.Fly()
}

func main() {
	var b Bird
	var fb FlyingBird

	b = Sparrow{}
	LetBirdWalk(b) // OK

	b = Penguin{}
	LetBirdWalk(b) // OK

	fb = Sparrow{}
	LetBirdFly(fb) // OK

	// Penguin does not implement FlyingBird, so this won't compile:
	// fb = Penguin{} // compile-time error -> safe
}

// What shows Liskov Substitution here?
// Penguin is only used where a generic Bird (that can walk) is expected; it doesn’t pretend it can fly.
// Sparrow can be used both as a Bird and as a FlyingBird with no surprises.
// Any value that satisfies an interface (Bird, FlyingBird) can safely replace another implementation without breaking expectations, which is exactly LSP.
// If you want, I can help you adapt this to your Lisko.go file in your SOLID examples structure.


// Yes, that’s very close — with one important nuance.
// Simple way to say LSP
// If some code works with an interface X, it should work the same with ANY implementation of X without:
// unexpected panics,
// “not supported” errors,
// or weird behavior that breaks the assumptions of that interface.
// So:
// Interface side: defines a contract (what callers can safely expect).
// Implementation side: must honor that contract fully – not require extra conditions, not return weaker results, not violate invariants.
// Example:
// If MessageSender says SendMessage(to, msg) will send a message:
// Every implementation (EmailSender, SMSSender, SlackSender) must really send (or reasonably fail) in a way callers can handle.
// None of them should silently ignore, panic, or behave totally differently in a way the interface never promised.