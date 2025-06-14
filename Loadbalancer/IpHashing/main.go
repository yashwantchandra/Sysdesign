package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)
type IPHash struct {
	servers []string
}
func NewIPHash(servers []string) *IPHash {
	return &IPHash{servers: servers}
}

func (lb *IPHash) GetNextServer(clientIP string) string {
	hasher := md5.New()
	hasher.Write([]byte(clientIP))
	hashValue := hex.EncodeToString(hasher.Sum(nil))

	// Convert a portion of the hash (e.g., first 8 characters) to int
	var hashInt int
	for i := 0; i < 8; i++ {
		hashInt = hashInt*16 + hexDigitToInt(hashValue[i])
	}

	index := hashInt % len(lb.servers)
	return lb.servers[index]
}

func hexDigitToInt(b byte) int {
	if b >= '0' && b <= '9' {
		return int(b - '0')
	}
	return int(b-'a') + 10
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	loadBalancer := NewIPHash(servers)

	clientIPs := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3", "192.168.0.4"}
	for _, ip := range clientIPs {
		server := loadBalancer.GetNextServer(ip)
		fmt.Printf("Client %s -> %s\n", ip, server)
	}
}