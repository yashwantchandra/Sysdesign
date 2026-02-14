package main

import "fmt"

// Subsystems
type CPU struct{}

func (c *CPU) Start() { fmt.Println("CPU started") }

type Memory struct{}

func (m *Memory) Load() { fmt.Println("Memory loaded") }

type Disk struct{}

func (d *Disk) Read() { fmt.Println("Disk read") }

// Facade: simple interface over subsystems.
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
