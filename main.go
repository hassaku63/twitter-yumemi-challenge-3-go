package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6}
	array2 := FlatMap(&array1, func(arr *[]int, v int, i int) []int {
		if v%2 == 0 {
			return []int{v, v}
		} else {
			return []int{v, v, v}
		}
	})

	fmt.Println(array2)
	// expect: [1 1 1 2 2 3 3 3 4 4 5 5 5 6 6]
}

func FlatMap(array *[]int, fn func(arr *[]int, value int, index int) []int) []int {
	result := make([]int, 0)
	for i, v := range *array {
		ret := fn(array, v, i)
		result = append(result, ret...)
	}
	return result
}
