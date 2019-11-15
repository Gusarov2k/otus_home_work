// Написать функцию Concat , которая
// получает несколько слайсов и склеивает
// их в один длинный. { {1, 2, 3}, {4, 5},
// {6, 7} } => {1, 2, 3, 4, 5, 6, 7}
package main

import (
	"fmt"
)

func concatSlace(slice ...[]int) []int {
	sumInts := 0
	for _, val := range slice {
		sumInts += len(val)
	}

	result := make([]int, 0, sumInts)
	for _, val := range slice {
		result = append(result, val...)
	}
	return result
}

func main() {
	first := []int{1, 2, 3}
	second := []int{4, 5}
	third := []int{6, 7}
	thirde := []int{6, 7}

	fmt.Println(concatSlace(first, second, third, thirde))
}
