package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LeastResponseTime struct {
	servers        []string
	responseTimes  []float64
}

func NewLeastResponseTime(servers []string) *LeastResponseTime {
	responseTimes := make([]float64, len(servers))
	return &LeastResponseTime{
		servers:       servers,
		responseTimes: responseTimes,
	}
}

func (lb *LeastResponseTime) GetNextServer() string {
	minIndex := 0
	minTime := lb.responseTimes[0]

	for i := 1; i < len(lb.responseTimes); i++ {
		if lb.responseTimes[i] < minTime {
			minTime = lb.responseTimes[i]
			minIndex = i
		}
	}
	return lb.servers[minIndex]
}

func (lb *LeastResponseTime) UpdateResponseTime(server string, responseTime float64) {
	for i, s := range lb.servers {
		if s == server {
			lb.responseTimes[i] = responseTime
			break
		}
	}
}

func simulateResponseTime() float64 {
	delay := rand.Float64()*(1.0-0.1) + 0.1 // Between 0.1 and 1.0
	time.Sleep(time.Duration(delay * float64(time.Second)))
	return delay
}

func main() {
	rand.Seed(time.Now().UnixNano())

	servers := []string{"Server1", "Server2", "Server3"}
	loadBalancer := NewLeastResponseTime(servers)

	for i := 0; i < 6; i++ {
		server := loadBalancer.GetNextServer()
		fmt.Printf("Request %d -> %s\n", i+1, server)

		responseTime := simulateResponseTime()
		loadBalancer.UpdateResponseTime(server, responseTime)

		fmt.Printf("Response Time: %.2fs\n", responseTime)
	}
}