// You can edit this code!
// Click here and start typing.

package main

import "fmt"

type Animal struct {
	Name string
	
}

type Bird interface {
	Fly() string
}

type crow struct {
}

func (c crow) Fly() string {
	return "crow can fly"
}

func (a Animal) Speak() string {
	return "Some sound"
}

type dog struct {
	Animal
}

func main() {
	var dg dog
	var cr crow

	fmt.Println(dg.Speak())
	fmt.Println(cr.Fly())

}
