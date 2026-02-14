package main

import (
	"fmt"
)

// Invoice has ONLY invoice data and business logic.
type Invoice struct {
	ID       int
	Customer string
	Amount   float64
	TaxRate  float64
}

func (i Invoice) TotalWithTax() float64 {
	return i.Amount + i.Amount*i.TaxRate
}

// InvoicePrinter has ONLY responsibility: displaying/printing invoices.
type InvoicePrinter struct{}

func (p InvoicePrinter) Print(i Invoice) {
	fmt.Printf("Invoice #%d\n", i.ID)
	fmt.Printf("Customer: %s\n", i.Customer)
	fmt.Printf("Amount: %.2f\n", i.Amount)
	fmt.Printf("Total with tax: %.2f\n", i.TotalWithTax())
}

// InvoiceRepository has ONLY responsibility: persistence (saving to DB/file).
type InvoiceRepository struct{}

func (r InvoiceRepository) Save(i Invoice) error {
	// Imagine this saves to DB or file.
	fmt.Printf("Saving invoice #%d for %s to database...\n", i.ID, i.Customer)
	return nil
}

func main() {
	invoice := Invoice{
		ID:       1,
		Customer: "Alice",
		Amount:   100.0,
		TaxRate:  0.18,
	}

	printer := InvoicePrinter{}
	repo := InvoiceRepository{}

	printer.Print(invoice)
	if err := repo.Save(invoice); err != nil {
		fmt.Println("Error saving invoice:", err)
	}
}