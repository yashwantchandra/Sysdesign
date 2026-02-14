package main

import (
	"fmt"
	"sync"
)

// Singleton ensures a type has only one instance and provides global access.
var (
	instance *Database
	once     sync.Once
)

type Database struct {
	name string
}

func GetDatabase() *Database {
	once.Do(func() {
		instance = &Database{name: "main-db"}
	})
	return instance
}

func (d *Database) Query(sql string) {
	fmt.Printf("[%s] Executing: %s\n", d.name, sql)
}

func main() {
	db1 := GetDatabase()
	db2 := GetDatabase()

	db1.Query("SELECT * FROM users")
	db2.Query("SELECT * FROM orders")

	fmt.Println("Same instance?", db1 == db2) // true
}
