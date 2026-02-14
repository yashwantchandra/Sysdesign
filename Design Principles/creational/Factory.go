package main

import "fmt"

// Product interface - what the factory creates.
type Vehicle interface {
	Drive() string
}

type Car struct{}

func (c Car) Drive() string { return "Driving a car" }

type Bike struct{}

func (b Bike) Drive() string { return "Riding a bike" }

// Factory: creates the right product based on input.
func NewVehicle(vehicleType string) (Vehicle, error) {
	switch vehicleType {
	case "car":
		return Car{}, nil
	case "bike":
		return Bike{}, nil
	default:
		return nil, fmt.Errorf("unknown vehicle type: %s", vehicleType)
	}
}

func main() {
	v1, _ := NewVehicle("car")
	v2, _ := NewVehicle("bike")

	fmt.Println(v1.Drive())
	fmt.Println(v2.Drive())
}
