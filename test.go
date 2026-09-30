package main

import "fmt"

func testBalance() {
	num2 := "123"

	for i := range num2 {
		fmt.Println(i)
	}
}