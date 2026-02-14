package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct {
	on bool
}

func (l *Light) TurnOn() {
	l.on = true
	fmt.Println("Light on")
}

func (l *Light) TurnOff() {
	l.on = false
	fmt.Println("Light off")
}

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

type Remote struct {
	cmd Command
}

func (r *Remote) SetCommand(c Command) {
	r.cmd = c
}

func (r *Remote) Press() {
	r.cmd.Execute()
}

func main() {
	light := &Light{}
	remote := &Remote{}

	remote.SetCommand(LightOnCommand{light: light})
	remote.Press()

	remote.SetCommand(LightOffCommand{light: light})
	remote.Press()
}
