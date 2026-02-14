package main

import "fmt"

type State interface {
	Handle(ctx *Context)
}

type Context struct {
	state State
}

func (c *Context) SetState(s State) {
	c.state = s
}

func (c *Context) Request() {
	c.state.Handle(c)
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
	ctx.Request()
	ctx.Request()
	ctx.Request()
}
