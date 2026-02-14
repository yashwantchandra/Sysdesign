# Design Patterns — Explained Simply

This guide explains each pattern in plain language, with a real-world analogy and when to use it.

---

## What is a design pattern?

A **design pattern** is a reusable way to solve a common problem in code.  
It’s not a library or a keyword — it’s an **idea** you implement (e.g. “have one shared instance” or “wrap an object to add behavior”).  
Learning patterns helps you name solutions and communicate with other developers.

---

## SOLID principles (the “why” behind good design)

These are **principles**, not patterns. They guide how you structure code.

| Principle | In one sentence |
|-----------|------------------|
| **Single Responsibility** | One type does one job. If it does two jobs, split it. |
| **Open/Closed** | Add new behavior by adding new code (e.g. new types), not by editing existing code. |
| **Liskov Substitution** | If something implements an interface, you can swap it with any other implementation without breaking callers. |
| **Interface Segregation** | Prefer small, focused interfaces. Don’t force a type to implement methods it can’t really support. |
| **Dependency Inversion** | Depend on abstractions (interfaces), not concrete types. High-level logic shouldn’t know about “MySQL” or “Redis”. |

---

## Creational patterns (how we create objects)

### 1. Singleton

**Idea:** Exactly **one instance** of a type in the whole program. Everyone uses that same instance.

**Real-world:** The company’s main database connection, or a single config object everyone reads from.

**When to use:** When having multiple instances would be wrong (e.g. two “main” DB connections) or wasteful (e.g. one config loader is enough).

**In code:** A function like `GetDatabase()` that creates the instance once (e.g. with `sync.Once`) and returns the same pointer every time.

---

### 2. Factory (Factory Method)

**Idea:** You ask for “a vehicle” or “a payment method” **by name or type**, and you get the right concrete object. The code that creates the object is in one place (the “factory”).

**Real-world:** “I want a car” → you get a car. “I want a bike” → you get a bike. You don’t build them yourself; the factory gives you the right thing.

**When to use:** When you don’t want callers to know the exact type (e.g. `Car` vs `Bike`) or the creation steps. Adding a new kind (e.g. “truck”) means changing only the factory.

**In code:** A function like `NewVehicle("car")` that returns a `Vehicle` interface. Callers use `Vehicle`; they don’t care if it’s a car or a bike.

---

### 3. Builder

**Idea:** Building a complex object **step by step** with a fluent API (e.g. `.SetWalls("brick").SetDoors(2).Build()`), instead of one huge constructor with many arguments.

**Real-world:** Ordering food: “base + cheese + tomato + olives” — you add options one by one, then get the final pizza.

**When to use:** When an object has many optional parts (walls, doors, roof, etc.) and you want readable, flexible construction without a constructor with 10 parameters.

**In code:** A `HouseBuilder` with methods like `SetWalls()`, `SetDoors()`, `Build()`. Each method returns the builder so you can chain: `builder.SetWalls("brick").SetDoors(2).Build()`.

---

## Structural patterns (how we compose and connect objects)

### 4. Adapter

**Idea:** You have an **existing type** that does what you need, but its method names or signatures don’t match what your code expects. The **adapter** wraps that type and exposes the interface your code expects.

**Real-world:** A power adapter: your device expects 5V; the wall gives 220V. The adapter in between makes them compatible.

**When to use:** When you’re integrating a library or legacy code that has a different API (e.g. `PrintLegacy(text)`) and you want your code to only use one interface (e.g. `Print(s)`).

**In code:** Your code depends on `Printer` with `Print(s)`. The legacy type has `PrintLegacy(text)`. `PrinterAdapter` wraps the legacy type and implements `Print(s)` by calling `PrintLegacy(s)`.

---

### 5. Decorator

**Idea:** You have a **base object** (e.g. “coffee”). You want to **add behavior or data** without changing the base. So you **wrap** it in a “decorator” that implements the same interface and adds something (e.g. “with milk”, “with sugar”). You can wrap multiple times.

**Real-world:** Coffee → add milk → add sugar. Same cup of coffee, but each wrapper adds one thing. The base stays “simple coffee”.

**When to use:** When you want to add optional behavior (logging, caching, extra fields) without subclassing or modifying the original type. Good when combinations grow (coffee, coffee+milk, coffee+sugar, coffee+milk+sugar).

**In code:** `SimpleCoffee` implements `Coffee`. `MilkDecorator` holds a `Coffee` and implements `Coffee` by delegating to it and adding milk (e.g. in `Cost()` and `Description()`). You build: `MilkDecorator{SugarDecorator{SimpleCoffee{}}}`.

---

### 6. Facade

**Idea:** A **single, simple interface** in front of several subsystems. Callers use one or two methods (e.g. “boot the computer”) instead of knowing about CPU, memory, disk, etc.

**Real-world:** Starting a car: you turn the key. You don’t operate fuel pump, ignition, and starter separately — the “facade” (turning the key) does that for you.

**When to use:** When you have many types or steps that usually go together, and you want to hide that complexity behind a simple API.

**In code:** `Computer` has `Boot()`. Inside `Boot()`, it calls `cpu.Start()`, `memory.Load()`, `disk.Read()`. Callers only call `computer.Boot()`.

---

### 7. Proxy

**Idea:** A **stand-in** for the real object. The proxy has the same interface as the real object, but it controls **when** and **how** the real object is used (e.g. create it only when needed, add logging, check permissions).

**Real-world:** A receptionist in front of a busy executive. You talk to the receptionist; they decide when to forward the call and what to say.

**When to use:** When you want lazy creation (create the heavy object only when first used), access control, logging, or a local stand-in for a remote service.

**In code:** `Proxy` implements the same interface as `RealSubject`. When `Request()` is called, the proxy may create `RealSubject` the first time, then forward the call. Callers use the proxy and don’t know the difference.

---

## Behavioral patterns (how objects interact and who does what)

### 8. Strategy

**Idea:** You have **one kind of action** (e.g. “pay”) but **multiple ways to do it** (credit card, PayPal, crypto). Instead of big if/else blocks, you make each way a separate type (a “strategy”) and plug the right one into the client.

**Real-world:** Choosing a route: “by car”, “by bike”, “by bus”. The goal is “go from A to B”; the strategy is how you get there.

**When to use:** When you have several interchangeable algorithms or behaviors and you want to add or change them without editing the code that uses them (e.g. payment methods, sort orders, compression algorithms).

**In code:** `PaymentStrategy` interface with `Pay(amount)`. `CreditCard` and `PayPal` implement it. `Cart` holds a `PaymentStrategy` and calls `strategy.Pay(amount)` at checkout. To add a new method, add a new type that implements the interface.

---

### 9. Observer

**Idea:** One object (the **subject**) has a list of **observers** (listeners). When something important happens, the subject **notifies** all observers. They can react (update UI, send email, log, etc.) without the subject knowing their details.

**Real-world:** News channel and subscribers. When news breaks, the channel broadcasts; each subscriber gets the update in their own way (TV, app, email).

**When to use:** When one thing (e.g. “order placed”, “price changed”) must trigger many independent reactions (email, SMS, analytics, logging) and you don’t want the core logic to depend on all of them.

**In code:** `Subject` has `Attach(observer)` and `Notify(msg)`. Each observer implements `Update(msg)`. When the subject’s state changes, it calls `Notify`; each observer’s `Update` runs.

---

### 10. Command

**Idea:** Turn a **request** or **action** into an **object**. That object has a single method like `Execute()`. So “turn on the light” becomes a `LightOnCommand` object. You can pass it around, queue it, undo it, or log it.

**Real-world:** Remote control: each button is a “command”. Pressing the button runs that command. The remote doesn’t know how the TV works; it just runs the command.

**When to use:** When you want to queue actions, support undo/redo, log every action, or let different parts of the system trigger the same action without knowing the details (e.g. UI button and keyboard shortcut both “save”).

**In code:** `Command` interface with `Execute()`. `LightOnCommand` holds a reference to the light and in `Execute()` calls `light.TurnOn()`. A `Remote` (or menu, or queue) holds a `Command` and calls `Execute()` when the user presses a button.

---

### 11. State

**Idea:** An object’s **behavior** depends on its internal **state** (e.g. “idle”, “running”, “paused”). Instead of lots of if/else on that state, you give each state its own type. The context holds the current state and delegates to it; the state can switch the context to another state.

**Real-world:** A traffic light: “green” → on timeout → “yellow” → “red” → “green”. Each state knows what comes next; the light just delegates to the current state.

**When to use:** When you have clear states and the behavior (and/or next state) is different in each state. Especially when adding a new state would otherwise mean touching many if/else branches.

**In code:** `State` interface with `Handle(ctx)`. `Context` holds the current `State`. When `Request()` is called, it calls `state.Handle(ctx)`. In `Handle`, the state may do work and then call `ctx.SetState(NextState{})`. Each state type encapsulates one state’s behavior and transitions.

---

## Quick reference

| Pattern   | One-line idea |
|----------|----------------|
| Singleton | One shared instance for the whole app. |
| Factory   | “Give me X” → one place creates and returns the right type. |
| Builder   | Build complex object step by step with a fluent API. |
| Adapter   | Wrap an existing type so it matches the interface you need. |
| Decorator | Wrap an object to add behavior without changing the original. |
| Facade    | One simple API in front of many subsystems. |
| Proxy     | Stand-in that controls access to the real object. |
| Strategy  | Plug in one of several interchangeable algorithms. |
| Observer  | One subject notifies many listeners when something happens. |
| Command   | Encapsulate an action as an object (execute, queue, undo). |
| State     | Behavior and transitions live in state objects; context delegates. |

Use this file together with the example code in `solid/`, `creational/`, `structural/`, and `behavioral/` to see each idea in real Go code.
