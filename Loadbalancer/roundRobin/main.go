package main
import (
	"fmt"
)
type RoundRobin struct {
	servers      []string
	currentIndex int
}
func NewRoundRobin(servers []string) *RoundRobin {
	return &RoundRobin{
		servers:      servers,
		currentIndex: -1,
	}
}

func (rr *RoundRobin) GetNextServer() string {
	rr.currentIndex = (rr.currentIndex + 1) % len(rr.servers)
	return rr.servers[rr.currentIndex]
}
func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	loadBalancer := NewRoundRobin(servers)
	for i := 0; i < 6; i++ {
		server := loadBalancer.GetNextServer()
		fmt.Printf("Request %d -> %s\n", i+1, server)
	}
}