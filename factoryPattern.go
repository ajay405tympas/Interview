// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type vehicle interface {
	getName() string
	getPower() int
}

type car struct {
	name  string
	power int
}

type truck struct {
	name  string
	power int
}

func (c *car) getName() string {
	return c.name
}

func (c *car) getPower() int {
	return c.power
}

func (t *truck) getName() string {
	return t.name
}

func (t *truck) getPower() int {
	return t.power
}

func GetVehicle(vehicleType string) (vehicle, error) {
	if vehicleType == "car" {
		return &car{name: "Benz ", power: 100}, nil
	}

	if vehicleType == "truck" {
		return &truck{name: "Leyland", power: 1500}, nil
	}
	return nil, fmt.Errorf("wrong vehicle type passed")
}

func printDetails(v vehicle) {
	fmt.Printf("Vechicle: %s, Power: %d\n", v.getName(), v.getPower())
}

func main() {

	car, _ := GetVehicle("car")
	truck, _ := GetVehicle("truck")

	printDetails(car)
	printDetails(truck)
	fmt.Println("Hello, 世界")
}
