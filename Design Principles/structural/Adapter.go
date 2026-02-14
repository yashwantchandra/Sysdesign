package main

import "fmt"

// Target: what our client expects.
type Printer interface {
	Print(s string)
}

// Adaptee: existing type with a different method name.
type LegacyPrinter struct{}

func (p *LegacyPrinter) PrintLegacy(text string) {
	fmt.Println("Legacy:", text)
}

// Adapter: wraps Adaptee and implements Target.
type PrinterAdapter struct {
	legacy *LegacyPrinter
}

func (a *PrinterAdapter) Print(s string) {
	a.legacy.PrintLegacy(s)
}

func main() {
	var p Printer = &PrinterAdapter{legacy: &LegacyPrinter{}}
	p.Print("Hello") // Client uses Print; adapter forwards to PrintLegacy
}
