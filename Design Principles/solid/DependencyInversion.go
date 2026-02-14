Dependency Inversion Principle (DIP) in simple words
Idea:
High-level modules (business logic) should not depend on low-level modules (e.g., DB, HTTP, file).
Both should depend on abstractions (interfaces).
Abstractions should not depend on details; details depend on abstractions.


Violating example (bad design)
High-level OrderService depends directly on a concrete MySQL implementation
package main

import "fmt"

// Low-level concrete dependency.
type MySQLOrderRepository struct{}

func (r MySQLOrderRepository) SaveOrder(orderID string) error {
	fmt.Println("Saving order to MySQL:", orderID)
	return nil
}

// High-level business logic directly depends on MySQLOrderRepository (detail).
type OrderService struct {
	repo MySQLOrderRepository
}

func NewOrderService() *OrderService {
	// Hard-coded dependency: can't easily swap DB or mock in tests.
	return &OrderService{
		repo: MySQLOrderRepository{},
	}
}

func (s *OrderService) PlaceOrder(orderID string) error {
	// business logic...
	fmt.Println("Placing order:", orderID)
	return s.repo.SaveOrder(orderID)
}

func main() {
	service := NewOrderService()
	_ = service.PlaceOrder("order-123")
}

Problems:
OrderService is tightly coupled to MySQLOrderRepository.
Hard to test (you always hit “MySQL” behavior).
Hard to switch to another storage (Postgres, in-memory, etc.) without editing OrderService.



DIP-compliant example (good design)
We introduce an abstraction OrderRepository and make both high-level and low-level depend on it.

package main

import "fmt"

// Abstraction that both high-level and low-level depend on.
type OrderRepository interface {
	SaveOrder(orderID string) error
}

// Low-level: concrete MySQL implementation depends on the abstraction.
type MySQLOrderRepository struct{}

func (r MySQLOrderRepository) SaveOrder(orderID string) error {
	fmt.Println("Saving order to MySQL:", orderID)
	return nil
}

// Another low-level: in-memory implementation (for tests or simple cases).
type InMemoryOrderRepository struct {
	data []string
}

func (r *InMemoryOrderRepository) SaveOrder(orderID string) error {
	r.data = append(r.data, orderID)
	fmt.Println("Saving order in memory:", orderID)
	return nil
}

// High-level business logic depends on the abstraction, not on concrete DB.
type OrderService struct {
	repo OrderRepository
}

// Dependency is injected (constructor injection).
func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) PlaceOrder(orderID string) error {
	// business logic...
	fmt.Println("Placing order:", orderID)
	return s.repo.SaveOrder(orderID)
}

func main() {
	// Production: use MySQL
	mysqlRepo := MySQLOrderRepository{}
	orderServiceMySQL := NewOrderService(mysqlRepo)
	_ = orderServiceMySQL.PlaceOrder("order-123")

	// Tests / dev: use in-memory
	memRepo := &InMemoryOrderRepository{}
	orderServiceMem := NewOrderService(memRepo)
	_ = orderServiceMem.PlaceOrder("order-456")
}


What shows Dependency Inversion here?
High-level OrderService depends on OrderRepository (an interface), not on MySQL directly.
Low-level details (MySQLOrderRepository, InMemoryOrderRepository) also depend on the same abstraction by implementing it.
Swapping implementations or testing with a fake repo is just changing what you inject into NewOrderService, without modifying OrderService itself.