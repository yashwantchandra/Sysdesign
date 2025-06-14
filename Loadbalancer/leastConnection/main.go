package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LeastConnections struct {
	servers map[string]int
}

func NewLeastConnections(serverList []string) *LeastConnections {
	servers := make(map[string]int)
	for _, server := range serverList {
		servers[server] = 0
	}
	return &LeastConnections{servers: servers}
}

func (lb *LeastConnections) GetNextServer() string {
	// Find the minimum number of connections
	minConnections := -1
	for _, conn := range lb.servers {
		if minConnections == -1 || conn < minConnections {
			minConnections = conn
		}
	}

	// Get all servers with the minimum number of connections
	var leastLoaded []string
	for server, conn := range lb.servers {
		if conn == minConnections {
			leastLoaded = append(leastLoaded, server)
		}
	}

	// Select one randomly from the least loaded
	rand.Seed(time.Now().UnixNano())
	selected := leastLoaded[rand.Intn(len(leastLoaded))]
	lb.servers[selected] += 1
	return selected
}

func (lb *LeastConnections) ReleaseConnection(server string) {
	if lb.servers[server] > 0 {
		lb.servers[server] -= 1
	}
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	loadBalancer := NewLeastConnections(servers)

	for i := 0; i < 6; i++ {
		server := loadBalancer.GetNextServer()
		fmt.Printf("Request %d -> %s\n", i+1, server)
		loadBalancer.ReleaseConnection(server)
	}
}