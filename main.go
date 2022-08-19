package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6}
	array2 := array1 /* TODO */

	fmt.Println(array2)
	// expect: [1 1 1 2 2 3 3 3 4 4 5 5 5 6 6]
}
