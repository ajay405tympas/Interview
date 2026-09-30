// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Personx struct {
	Namex string
	Agex  int
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case Personx:
		fmt.Println("Person:", v.Namex, v.Agex)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	p := Person{Name: "Ajay", Age: 25}
	checkType(p)

	var i interface{} = "hello"

	str:=i.(string)

	//checkType(i)

	fmt.Println("str:", str)

	fmt.Println("Hello, 世界")
}
