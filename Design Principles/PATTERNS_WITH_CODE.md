# Design Patterns — Description + Code

Each section has: **what it is**, **when to use it**, and **full code with comments**.  
Run any example: `go run solid/singleResponsibility.go` (or the path to the file).

---

# SOLID Principles

---

## 1. Single Responsibility (SRP)

**Description:** One type has **one reason to change**. If it does two jobs (e.g. invoice logic + printing + saving), split into separate types. That way, changing how we print doesn’t touch how we save.

**When to use:** When a struct or package is doing more than one clear job, or when changes in one area keep touching the same file.

**Code:** (`solid/singleResponsibility.go`)

```go
package main

import "fmt"

// Invoice: ONLY data + business logic (e.g. total with tax).
type Invoice struct {
	ID       int
	Customer string
	Amount   float64
	TaxRate  float64
}

func (i Invoice) TotalWithTax() float64 {
	return i.Amount + i.Amount*i.TaxRate
}

// InvoicePrinter: ONLY printing/display. No saving, no business rules.
type InvoicePrinter struct{}

func (p InvoicePrinter) Print(i Invoice) {
	fmt.Printf("Invoice #%d\n", i.ID)
	fmt.Printf("Customer: %s\n", i.Customer)
	fmt.Printf("Amount: %.2f\n", i.Amount)
	fmt.Printf("Total with tax: %.2f\n", i.TotalWithTax())
}

// InvoiceRepository: ONLY persistence. No printing, no tax calculation.
type InvoiceRepository struct{}

func (r InvoiceRepository) Save(i Invoice) error {
	fmt.Printf("Saving invoice #%d for %s to database...\n", i.ID, i.Customer)
	return nil
}

func main() {
	invoice := Invoice{ID: 1, Customer: "Alice", Amount: 100.0, TaxRate: 0.18}
	printer := InvoicePrinter{}
	repo := InvoiceRepository{}

	printer.Print(invoice)
	_ = repo.Save(invoice)
}
```

**Takeaway:** Three types, three jobs. Change print format → touch only `InvoicePrinter`. Change storage → only `InvoiceRepository`.

---

## 2. Open/Closed (OCP)

**Description:** **Open for extension, closed for modification.** You add new behavior by adding new types (e.g. new senders) that implement an interface, without changing existing types (e.g. `Notifier`).

**When to use:** When you expect to add new variants (notification channels, payment methods, exporters) and want to avoid editing stable, tested code.

**Code:** (`solid/OpenClose.go`)

```go
package main

import "fmt"

// Abstraction: Notifier depends on this, not on EmailSender/SMSSender.
type MessageSender interface {
	SendMessage(to, message string) error
}

type EmailSender struct{}

func (e EmailSender) SendMessage(to, message string) error {
	fmt.Printf("Sending EMAIL to %s: %s\n", to, message)
	return nil
}

type SMSSender struct{}

func (s SMSSender) SendMessage(to, message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
	return nil
}

// Notifier is "closed for modification": it never changes when we add SlackSender, etc.
type Notifier struct {
	sender MessageSender
}

func NewNotifier(sender MessageSender) *Notifier {
	return &Notifier{sender: sender}
}

func (n *Notifier) NotifyUser(userID, message string) error {
	return n.sender.SendMessage(userID, message)
}

func main() {
	emailNotifier := NewNotifier(EmailSender{})
	emailNotifier.NotifyUser("alice@example.com", "Welcome!")

	smsNotifier := NewNotifier(SMSSender{})
	smsNotifier.NotifyUser("+1234567890", "OTP 1234")
}
```

**Takeaway:** To add Slack: create `SlackSender` implementing `MessageSender` and pass it to `NewNotifier`. No change inside `Notifier`.

---

## 3. Liskov Substitution (LSP)

**Description:** If a type implements an interface, **any use of that interface must work with that type** — no panics, no “not supported,” no broken assumptions. Subtypes must be substitutable for the base.

**When to use:** When defining interfaces: only put methods that every implementation can honestly support. If some implementations can’t do something, use a smaller or separate interface.

**Code:** (`solid/Lisko.go`)

```go
package main

import "fmt"

// General: every bird can walk.
type Bird interface {
	Walk()
}

// Only birds that can fly implement this. Penguin does NOT.
type FlyingBird interface {
	Bird
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Walk() { fmt.Println("Sparrow is walking") }
func (s Sparrow) Fly()  { fmt.Println("Sparrow is flying") }

type Penguin struct{}

func (p Penguin) Walk() { fmt.Println("Penguin is walking") }
// No Fly() — penguins don't implement FlyingBird, so no fake/panic.

func LetBirdWalk(b Bird)       { b.Walk() }
func LetBirdFly(fb FlyingBird) { fb.Fly() }

func main() {
	LetBirdWalk(Sparrow{})  // OK
	LetBirdWalk(Penguin{})  // OK
	LetBirdFly(Sparrow{})   // OK
	// LetBirdFly(Penguin{}) // compile error: Penguin doesn't implement FlyingBird
}
```

**Takeaway:** Callers that need “something that can fly” use `FlyingBird`; they can substitute any implementation (e.g. Sparrow) without surprises.

---

## 4. Interface Segregation (ISP)

**Description:** **Many small interfaces** over one fat interface. Clients should not depend on methods they don’t use. Implementations shouldn’t be forced to provide stubs or panics for unsupported operations.

**When to use:** When an interface has many methods and some types only support a subset. Split into Printer, Scanner, Faxer so a simple printer only implements Printer.

**Code:** (`solid/InterfaceSegregation.go`)

```go
package main

import "fmt"

// Small interfaces: clients depend only on what they use.
type Printer interface {
	Print(document string)
}
type Scanner interface {
	Scan(document string)
}
type Faxer interface {
	Fax(document string)
}

type OfficePrinter struct{}

func (p OfficePrinter) Print(d string)  { fmt.Println("Printing:", d) }
func (p OfficePrinter) Scan(d string)  { fmt.Println("Scanning:", d) }
func (p OfficePrinter) Fax(d string)   { fmt.Println("Faxing:", d) }

// SimplePrinter only implements Printer — no Scan/Fax, no panics.
type SimplePrinter struct{}

func (p SimplePrinter) Print(document string) {
	fmt.Println("Printing:", document)
}

func PrintDocument(p Printer, doc string) { p.Print(doc) }
func ScanDocument(s Scanner, doc string)  { s.Scan(doc) }

func main() {
	PrintDocument(SimplePrinter{}, "Invoice")

	office := OfficePrinter{}
	PrintDocument(office, "Report")
	ScanDocument(office, "Contract")
}
```

**Takeaway:** Callers that only need printing use `Printer`; they don’t depend on Scan or Fax. SimplePrinter doesn’t implement methods it can’t support.

---

## 5. Dependency Inversion (DIP)

**Description:** **Depend on abstractions, not concretions.** High-level logic (e.g. OrderService) should not depend on low-level details (e.g. MySQL). Both depend on an interface (e.g. OrderRepository). The concrete implementation is injected (e.g. in tests you inject an in-memory repo).

**When to use:** Whenever high-level code would otherwise import a specific DB, HTTP client, or file system. Introduce an interface and inject the implementation so you can test and swap without changing the high-level code.

**Code:** (`solid/DependencyInversion.go`)

```go
package main

import "fmt"

// Abstraction: high-level and low-level both depend on this.
type OrderRepository interface {
	SaveOrder(orderID string) error
}

type MySQLOrderRepository struct{}

func (r MySQLOrderRepository) SaveOrder(orderID string) error {
	fmt.Println("Saving order to MySQL:", orderID)
	return nil
}

type InMemoryOrderRepository struct {
	data []string
}

func (r *InMemoryOrderRepository) SaveOrder(orderID string) error {
	r.data = append(r.data, orderID)
	fmt.Println("Saving order in memory:", orderID)
	return nil
}

// OrderService depends on OrderRepository (abstraction), not MySQL.
type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) PlaceOrder(orderID string) error {
	fmt.Println("Placing order:", orderID)
	return s.repo.SaveOrder(orderID)
}

func main() {
	// Production: inject MySQL
	svc1 := NewOrderService(MySQLOrderRepository{})
	_ = svc1.PlaceOrder("order-123")

	// Tests: inject in-memory
	svc2 := NewOrderService(&InMemoryOrderRepository{})
	_ = svc2.PlaceOrder("order-456")
}
```

**Takeaway:** `OrderService` never imports MySQL. You swap storage by passing a different `OrderRepository` into `NewOrderService`.

---

# Creational Patterns

---

## 6. Singleton

**Description:** Ensures **exactly one instance** of a type in the program and provides a single access point (e.g. `GetDatabase()`). Often used for shared resources (DB connection, config, logger).

**When to use:** When multiple instances would be wrong or wasteful (one config, one connection pool). Use sparingly; it can make testing harder.

**Code:** (`creational/Singleton.go`)

```go
package main

import (
	"fmt"
	"sync"
)

var (
	instance *Database
	once     sync.Once  // ensures creation runs only once
)

type Database struct {
	name string
}

func GetDatabase() *Database {
	once.Do(func() {
		instance = &Database{name: "main-db"}
	})
	return instance
}

func (d *Database) Query(sql string) {
	fmt.Printf("[%s] Executing: %s\n", d.name, sql)
}

func main() {
	db1 := GetDatabase()
	db2 := GetDatabase()

	db1.Query("SELECT * FROM users")
	db2.Query("SELECT * FROM orders")

	fmt.Println("Same instance?", db1 == db2)  // true
}
```

**Takeaway:** Every call to `GetDatabase()` returns the same `*Database`. Safe for concurrent use thanks to `sync.Once`.

---

## 7. Factory (Factory Method)

**Description:** A **factory function** creates the right concrete type based on input (e.g. `"car"` vs `"bike"`) and returns an interface. Callers use the interface and don’t depend on concrete structs.

**When to use:** When creation logic is non-trivial or you want to centralize it so adding a new product type (e.g. "truck") only touches the factory.

**Code:** (`creational/Factory.go`)

```go
package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct{}

func (c Car) Drive() string { return "Driving a car" }

type Bike struct{}

func (b Bike) Drive() string { return "Riding a bike" }

// Factory: one place that knows how to create Car vs Bike.
func NewVehicle(vehicleType string) (Vehicle, error) {
	switch vehicleType {
	case "car":
		return Car{}, nil
	case "bike":
		return Bike{}, nil
	default:
		return nil, fmt.Errorf("unknown vehicle type: %s", vehicleType)
	}
}

func main() {
	v1, _ := NewVehicle("car")
	v2, _ := NewVehicle("bike")

	fmt.Println(v1.Drive())
	fmt.Println(v2.Drive())
}
```

**Takeaway:** Callers get a `Vehicle` and call `Drive()`. They don’t need to know about `Car` or `Bike`; the factory decides.

---

## 8. Builder

**Description:** Builds a complex object **step by step** via a fluent API (e.g. `SetWalls("brick").SetDoors(2).Build()`). Avoids huge constructor argument lists and makes optional parts clear.

**When to use:** When an object has many fields or optional components and you want readable, flexible construction.

**Code:** (`creational/Builder.go`)

```go
package main

import "fmt"

type House struct {
	Walls   string
	Doors   int
	Windows int
	Roof    string
}

type HouseBuilder struct {
	house House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

// Each setter returns the builder so calls can be chained.
func (b *HouseBuilder) SetWalls(w string) *HouseBuilder {
	b.house.Walls = w
	return b
}

func (b *HouseBuilder) SetDoors(n int) *HouseBuilder {
	b.house.Doors = n
	return b
}

func (b *HouseBuilder) SetWindows(n int) *HouseBuilder {
	b.house.Windows = n
	return b
}

func (b *HouseBuilder) SetRoof(r string) *HouseBuilder {
	b.house.Roof = r
	return b
}

func (b *HouseBuilder) Build() House {
	return b.house
}

func main() {
	house := NewHouseBuilder().
		SetWalls("brick").
		SetDoors(2).
		SetWindows(4).
		SetRoof("tile").
		Build()

	fmt.Printf("House: %+v\n", house)
}
```

**Takeaway:** Construction is readable and optional steps can be omitted. `Build()` returns the final `House`.

---

# Structural Patterns

---

## 9. Adapter

**Description:** **Adapts** an existing type to match an interface your code expects. The adapter wraps the “adaptee” and implements the target interface by delegating to the adaptee’s methods (possibly with a different name or signature).

**When to use:** When integrating a library or legacy code whose API doesn’t match the interface you want to depend on.

**Code:** (`structural/Adapter.go`)

```go
package main

import "fmt"

// Target: what our code expects.
type Printer interface {
	Print(s string)
}

// Adaptee: existing type with different method name.
type LegacyPrinter struct{}

func (p *LegacyPrinter) PrintLegacy(text string) {
	fmt.Println("Legacy:", text)
}

// Adapter: implements Printer by wrapping LegacyPrinter and forwarding.
type PrinterAdapter struct {
	legacy *LegacyPrinter
}

func (a *PrinterAdapter) Print(s string) {
	a.legacy.PrintLegacy(s)
}

func main() {
	var p Printer = &PrinterAdapter{legacy: &LegacyPrinter{}}
	p.Print("Hello")  // client calls Print; adapter calls PrintLegacy
}
```

**Takeaway:** Client code depends only on `Printer`. The adapter hides the mismatch between `Print(s)` and `PrintLegacy(text)`.

---

## 10. Decorator

**Description:** **Wraps** an object to add behavior (or state) without changing the original type. The decorator implements the same interface as the wrapped object and delegates to it, adding something (e.g. extra cost, logging). You can wrap multiple times (e.g. coffee → milk → sugar).

**When to use:** When you want to add optional behavior (logging, caching, extra fields) without subclassing or modifying the core type.

**Code:** (`structural/Decorator.go`)

```go
package main

import "fmt"

type Coffee interface {
	Cost() int
	Description() string
}

type SimpleCoffee struct{}

func (c SimpleCoffee) Cost() int            { return 10 }
func (c SimpleCoffee) Description() string { return "Simple coffee" }

// Decorator: wraps a Coffee and adds milk.
type MilkDecorator struct {
	coffee Coffee
}

func (m MilkDecorator) Cost() int {
	return m.coffee.Cost() + 2
}

func (m MilkDecorator) Description() string {
	return m.coffee.Description() + ", milk"
}

// Another decorator: adds sugar.
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
	// Simple coffee, milk, sugar -> 13 cents
}
```

**Takeaway:** Each decorator implements `Coffee` and holds another `Coffee`. You stack them; the outer decorator adds its part and delegates to the inner one.

---

## 11. Facade

**Description:** Provides a **single, simple interface** to a set of subsystems. Callers use one or two methods (e.g. `Boot()`) instead of knowing about CPU, memory, disk, etc.

**When to use:** When you have many types or steps that usually go together and you want to hide that complexity behind a simple API.

**Code:** (`structural/Facade.go`)

```go
package main

import "fmt"

type CPU struct{}

func (c *CPU) Start() { fmt.Println("CPU started") }

type Memory struct{}

func (m *Memory) Load() { fmt.Println("Memory loaded") }

type Disk struct{}

func (d *Disk) Read() { fmt.Println("Disk read") }

// Facade: one struct that owns subsystems and exposes a simple operation.
type Computer struct {
	cpu    *CPU
	memory *Memory
	disk   *Disk
}

func NewComputer() *Computer {
	return &Computer{
		cpu:    &CPU{},
		memory: &Memory{},
		disk:   &Disk{},
	}
}

func (c *Computer) Boot() {
	c.cpu.Start()
	c.memory.Load()
	c.disk.Read()
	fmt.Println("Computer ready")
}

func main() {
	pc := NewComputer()
	pc.Boot()
}
```

**Takeaway:** Callers only use `Computer` and `Boot()`. They don’t need to know about CPU, Memory, or Disk.

---

## 12. Proxy

**Description:** A **stand-in** for the real object. The proxy implements the same interface and controls access: e.g. lazy creation (create real object on first use), logging, or access control. Callers use the proxy as if it were the real thing.

**When to use:** When you want lazy initialization, access control, logging, or a local representative for a remote service.

**Code:** (`structural/Proxy.go`)

```go
package main

import "fmt"

type Subject interface {
	Request() string
}

type RealSubject struct{}

func (r *RealSubject) Request() string {
	return "real response"
}

// Proxy: same interface as RealSubject; creates real object only when needed.
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
	fmt.Println(sub.Request())  // creates RealSubject, then returns "real response"
	fmt.Println(sub.Request())  // reuses existing instance
}
```

**Takeaway:** Client talks to `Subject`. The proxy delays creating `RealSubject` until the first `Request()` and then forwards all calls.

---

# Behavioral Patterns

---

## 13. Strategy

**Description:** **Interchangeable algorithms** behind a common interface. The client holds a “strategy” (e.g. payment method) and delegates to it. You can add or change strategies without changing the client (e.g. Cart).

**When to use:** When you have several ways to do the same thing (payment, sorting, compression) and want to plug in one at a time without if/else in the client.

**Code:** (`behavioral/Strategy.go`)

```go
package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f with credit card", amount)
}

type PayPal struct{}

func (p PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f with PayPal", amount)
}

// Cart holds a strategy and delegates checkout to it.
type Cart struct {
	strategy PaymentStrategy
}

func (c *Cart) SetPayment(s PaymentStrategy) {
	c.strategy = s
}

func (c *Cart) Checkout(amount float64) string {
	return c.strategy.Pay(amount)
}

func main() {
	cart := &Cart{}

	cart.SetPayment(CreditCard{})
	fmt.Println(cart.Checkout(99.99))

	cart.SetPayment(PayPal{})
	fmt.Println(cart.Checkout(49.99))
}
```

**Takeaway:** `Cart` doesn’t know about credit cards or PayPal; it only calls `strategy.Pay(amount)`. New payment methods = new types implementing `PaymentStrategy`.

---

## 14. Observer

**Description:** One **subject** maintains a list of **observers** and notifies them when something happens. Observers implement a small interface (e.g. `Update(msg)`). Loose coupling: the subject doesn’t know what the observers do (email, log, UI).

**When to use:** When one event (e.g. “order placed”, “price changed”) should trigger many independent reactions without the core logic depending on all of them.

**Code:** (`behavioral/Observer.go`)

```go
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
```

**Takeaway:** Subject keeps a slice of `Observer` and calls `Update` on each when something happens. Add/remove observers without changing the subject’s logic.

---

## 15. Command

**Description:** Encapsulates a **request as an object** with a single method (e.g. `Execute()`). The invoker (e.g. remote, menu, queue) holds a command and calls `Execute()` without knowing the concrete action. Enables queuing, undo, logging.

**When to use:** When you want to queue actions, support undo/redo, or let different triggers (button, shortcut) run the same action without depending on the receiver.

**Code:** (`behavioral/Command.go`)

```go
package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct {
	on bool
}

func (l *Light) TurnOn()  { l.on = true; fmt.Println("Light on") }
func (l *Light) TurnOff() { l.on = false; fmt.Println("Light off") }

// Command object: holds receiver and calls the right method in Execute().
type LightOnCommand struct {
	light *Light
}

func (c LightOnCommand) Execute() {
	c.light.TurnOn()
}

type LightOffCommand struct {
	light *Light
}

func (c LightOffCommand) Execute() {
	c.light.TurnOff()
}

// Invoker: holds a command and runs it when "pressed".
type Remote struct {
	cmd Command
}

func (r *Remote) SetCommand(c Command) { r.cmd = c }
func (r *Remote) Press()               { r.cmd.Execute() }

func main() {
	light := &Light{}
	remote := &Remote{}

	remote.SetCommand(LightOnCommand{light: light})
	remote.Press()

	remote.SetCommand(LightOffCommand{light: light})
	remote.Press()
}
```

**Takeaway:** Remote doesn’t know about `Light`; it only has a `Command` and calls `Execute()`. You can swap commands (e.g. fan on/off) without changing the remote.

---

## 16. State

**Description:** An object’s **behavior** depends on its current **state**. Each state is a type implementing a small interface (e.g. `Handle(ctx)`). The context holds the current state and delegates to it; the state can switch the context to another state. Replaces big if/else on state.

**When to use:** When behavior and transitions differ a lot per state (e.g. idle vs running vs paused) and you want to add or change states without touching a central switch.

**Code:** (`behavioral/State.go`)

```go
package main

import "fmt"

type State interface {
	Handle(ctx *Context)
}

type Context struct {
	state State
}

func (c *Context) SetState(s State) { c.state = s }

func (c *Context) Request() {
	c.state.Handle(c)  // delegate to current state
}

type StateA struct{}

func (s StateA) Handle(ctx *Context) {
	fmt.Println("State A handling; switching to B")
	ctx.SetState(StateB{})
}

type StateB struct{}

func (s StateB) Handle(ctx *Context) {
	fmt.Println("State B handling; switching to A")
	ctx.SetState(StateA{})
}

func main() {
	ctx := &Context{state: StateA{}}
	ctx.Request()  // A handles, switches to B
	ctx.Request()  // B handles, switches to A
	ctx.Request()  // A again
}
```

**Takeaway:** Context delegates every request to the current `State`. Each state does its work and can transition by calling `ctx.SetState(...)`.

---

# Quick reference

| Name | One line |
|------|----------|
| **Single Responsibility** | One type, one job. |
| **Open/Closed** | Extend by new types; don’t modify existing ones. |
| **Liskov Substitution** | Implementations must honor the interface contract. |
| **Interface Segregation** | Small interfaces; no forced “not supported” methods. |
| **Dependency Inversion** | Depend on interfaces; inject concretions. |
| **Singleton** | One shared instance. |
| **Factory** | One place creates the right concrete type. |
| **Builder** | Fluent step-by-step construction. |
| **Adapter** | Wrap existing API to match target interface. |
| **Decorator** | Wrap to add behavior. |
| **Facade** | Simple API over many subsystems. |
| **Proxy** | Stand-in that controls access. |
| **Strategy** | Pluggable algorithm behind an interface. |
| **Observer** | Subject notifies many observers. |
| **Command** | Request as object; Execute() + invoker. |
| **State** | Behavior in state types; context delegates. |

Run from repo root, for example:

```bash
go run "Design Principles/solid/singleResponsibility.go"
go run "Design Principles/creational/Builder.go"
go run "Design Principles/behavioral/Observer.go"
```
