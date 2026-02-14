# Design Principles & Patterns

## SOLID Principles (`solid/`)

| Principle | File | Description |
|-----------|------|-------------|
| **S**ingle Responsibility | `singleResponsibility.go` | One type, one reason to change |
| **O**pen/Closed | `OpenClose.go` | Open for extension, closed for modification |
| **L**iskov Substitution | `Lisko.go` | Subtypes must be substitutable for their base types |
| **I**nterface Segregation | `InterfaceSegregation.go` | Many small interfaces over one fat interface |
| **D**ependency Inversion | `DependencyInversion.go` | Depend on abstractions, not concretions |

---

## Creational Patterns (`creational/`)

| Pattern | File | Description |
|---------|------|-------------|
| **Singleton** | `Singleton.go` | Single shared instance (e.g. DB connection) |
| **Factory** | `Factory.go` | Create objects without specifying exact type |
| **Builder** | `Builder.go` | Step-by-step construction of complex objects |

---

## Structural Patterns (`structural/`)

| Pattern | File | Description |
|---------|------|-------------|
| **Adapter** | `Adapter.go` | Make an existing type match a target interface |
| **Decorator** | `Decorator.go` | Add behavior by wrapping (e.g. coffee + milk + sugar) |
| **Facade** | `Facade.go` | Simple interface over many subsystems |
| **Proxy** | `Proxy.go` | Stand-in that controls access to the real object |

---

## Behavioral Patterns (`behavioral/`)

| Pattern | File | Description |
|---------|------|-------------|
| **Strategy** | `Strategy.go` | Interchangeable algorithms (e.g. payment methods) |
| **Observer** | `Observer.go` | Notify many listeners when subject changes |
| **Command** | `Command.go` | Encapsulate requests as objects (e.g. remote control) |
| **State** | `State.go` | Object behavior changes with internal state |

---

## Run examples

From this folder or any subfolder:

```bash
go run solid/singleResponsibility.go
go run creational/Singleton.go
go run structural/Adapter.go
go run behavioral/Strategy.go
```

Or from repo root:

```bash
go run "Design Principles/creational/Singleton.go"
```
