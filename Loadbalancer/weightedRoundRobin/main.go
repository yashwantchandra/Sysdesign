package main

import (
	"fmt"
)

type WeightedRoundRobin struct {
	servers        []string
	weights        []int
	currentIndex   int
	currentWeight  int
	maxWeight      int
}

func NewWeightedRoundRobin(servers []string, weights []int) *WeightedRoundRobin {
	return &WeightedRoundRobin{
		servers:       servers,
		weights:       weights,
		currentIndex:  -1,
		currentWeight: 0,
		maxWeight:     max(weights),
	}
}

func max(arr []int) int {
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func (w *WeightedRoundRobin) GetNextServer() string {
	for {
		w.currentIndex = (w.currentIndex + 1) % len(w.servers)
		if w.currentIndex == 0 {
			w.currentWeight--
			if w.currentWeight <= 0 {
				w.currentWeight = w.maxWeight
			}
		}
		if w.weights[w.currentIndex] >= w.currentWeight {
			return w.servers[w.currentIndex]
		}
	}
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	weights := []int{5, 1, 1}
	loadBalancer := NewWeightedRoundRobin(servers, weights)

	for i := 0; i < 7; i++ {
		server := loadBalancer.GetNextServer()
		fmt.Printf("Request %d -> %s\n", i+1, server)
	}
}