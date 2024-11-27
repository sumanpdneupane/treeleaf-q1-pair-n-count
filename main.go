package main

import "fmt"

func main() {
	array := []int{1, 9, 6, 4, 5}
	count, pairs := CountAndGetPairInversionFromArray(array)
	fmt.Printf("Total Inversions: %d\n", count)
	fmt.Printf("Inversion Pairs: %v\n", pairs)
}

func CountAndGetPairInversionFromArray(arr []int) (int, [][2]int) {
	count := 0
	var pairs [][2]int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
				pairs = append(pairs, [2]int{i, j})
			}
		}
	}

	return count, pairs
}
