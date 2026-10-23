// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{4, 5, 3, 8, 1, 2}

	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	fmt.Println(nums)

	str := []string{"Ajay", "xyz", "goat"}
	slices.SortFunc(str, func(a, b string) int {
		if a < b {
			return 1
		}
		if a > b {
			return -1
		}
		return 0

	})
	fmt.Println(str)

	str2 := []string{"Ajaydd", "go", "xyz", "goat"}
	slices.SortFunc(str2, func(a, b string) int {
		return len(a) - len(b)
	})
	fmt.Println(str2)

	type employee struct {
		name string
		age  int
	}

	str3 := []employee{
		{"Ajaydd", 5},
		{"go", 1},
		{"xyz", 3},
		{"goat", 4},
	}
	slices.SortFunc(str3, func(a, b employee) int {
		return a.age - b.age
	})
	fmt.Println(str3)

	str4 := []employee{
		{"Ajaydd", 5},
		{"mot", 1},
		{"xyz", 3},
		{"pgoat", 4},
	}
	slices.SortFunc(str4, func(a, b employee) int {
		if a.name < b.name {
			return -1
		}
		if a.name > b.name {
			return 1
		}
		return 0
	})
	fmt.Println(str4)

	fmt.Println("Hello, 世界")
}
