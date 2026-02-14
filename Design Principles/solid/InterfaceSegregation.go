Violating example (bad, “fat” interface):
package main

import "fmt"

// Fat interface: forces all implementers to support all operations.
type MultiFunctionPrinter interface {
	Print(document string)
	Scan(document string)
	Fax(document string)
}

// This machine can do everything -> OK (but interface is still fat).
type OfficePrinter struct{}

func (p OfficePrinter) Print(document string) {
	fmt.Println("Printing:", document)
}

func (p OfficePrinter) Scan(document string) {
	fmt.Println("Scanning:", document)
}

func (p OfficePrinter) Fax(document string) {
	fmt.Println("Faxing:", document)
}

// This machine can only print.
// ISP problem: it is forced to implement Scan and Fax anyway.
type SimplePrinter struct{}

func (p SimplePrinter) Print(document string) {
	fmt.Println("Printing:", document)
}

func (p SimplePrinter) Scan(document string) {
	// Not really supported…
	panic("Scan not supported on SimplePrinter")
}

func (p SimplePrinter) Fax(document string) {
	// Not really supported…
	panic("Fax not supported on SimplePrinter")
}

func main() {
	var device MultiFunctionPrinter

	device = OfficePrinter{}
	device.Print("Report") // OK

	device = SimplePrinter{}
	device.Print("Invoice") // OK
	// But if some code tries device.Scan("Doc"), it will panic.
}
What’s wrong?
SimplePrinter is forced to implement Scan and Fax which it cannot really support.
Callers that only need Print are still coupled to Scan and Fax.
This leads to panics or “not implemented” methods.
ISP-compliant design (good, small interfaces)
We split the fat interface into smaller, role-based interfaces.

// fixed:
package main

import "fmt"

// Small, focused interfaces.
type Printer interface {
	Print(document string)
}

type Scanner interface {
	Scan(document string)
}

type Faxer interface {
	Fax(document string)
}

// A full-featured device can implement multiple interfaces.
type OfficePrinter struct{}

func (p OfficePrinter) Print(document string) {
	fmt.Println("Printing:", document)
}

func (p OfficePrinter) Scan(document string) {
	fmt.Println("Scanning:", document)
}

func (p OfficePrinter) Fax(document string) {
	fmt.Println("Faxing:", document)
}

// SimplePrinter only needs to implement what it truly supports.
type SimplePrinter struct{}

func (p SimplePrinter) Print(document string) {
	fmt.Println("Printing:", document)
}

// Client that only needs printing depends only on Printer.
func PrintDocument(p Printer, doc string) {
	p.Print(doc)
}

// Another client that needs scanning only depends on Scanner.
func ScanDocument(s Scanner, doc string) {
	s.Scan(doc)
}

func main() {
	var printer Printer
	var scanner Scanner

	printer = SimplePrinter{}
	PrintDocument(printer, "Invoice") // OK, no unused methods

	office := OfficePrinter{}
	printer = office
	scanner = office

	PrintDocument(printer, "Report")
	ScanDocument(scanner, "Contract")
}

What shows Interface Segregation here?
Clients (PrintDocument, ScanDocument) depend only on what they use (Printer, Scanner).
Implementations (SimplePrinter, OfficePrinter) implement only the interfaces that match their capabilities.
No dummy / panic / “not supported” methods needed, and changes to scanning or faxing don’t affect code that only prints.