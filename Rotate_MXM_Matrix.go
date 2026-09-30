// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func rotateMatrix(m [][]int) {
	numRows := len(m)
	numCols := len(m[0])
	fmt.Println("Rows:", numRows)
	fmt.Println("Cols:", numCols)

	if numRows == numCols {
		numSwap := numRows / 2
		for i := 0; i < numSwap; i++ {
			m[i], m[numRows-(i+1)] = m[numRows-(i+1)], m[i]
		}
	} else {
		fmt.Println("Rotate Matrix possible only for square matrices")
		return
	}

	for i := 0; i < numRows; i++ {

		for j := 0; j < numCols; j++ {

			if i < j && i != j {
				m[i][j], m[j][i] = m[j][i], m[i][j]
			}
		}
	}

}

func main() {

	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	fmt.Println("Matrix before:")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(" ", matrix[i][j])
		}
		fmt.Println()
	}

	rotateMatrix(matrix)

	fmt.Println("Matrix after:")

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(" ", matrix[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Hello, 世界")
}
