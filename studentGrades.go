// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func gradeStudents(grades []int, len int) []int {

	result := make([]int, len)
	for i := range grades {
		if grades[i] > 37 {
			if grades[i]%5 >= 3 {
				result[i] = grades[i] + 5 - (grades[i] % 5)
			} else if grades[i]%5 < 3 {
				result[i] = grades[i]
			}
		} else {
			result[i] = grades[i]
			continue
		}
	}
	return result
}

func main() {
	grades := []int{24, 45, 68, 92}

	result := gradeStudents(grades, len(grades))

	fmt.Println("Print Result", result)

	fmt.Println("Hello, 世界")
}
