package main

import "fmt"

func main() {
	arr1 := []int{5, 1, 22, 25, 6, -1, 8, 10}
	arr2 := []int{1, 6, -1, 10}

	arr1InitIndex := 0
	arr2InitIndex := 0

	for {
		if len(arr1) <= arr1InitIndex || len(arr2) <= arr2InitIndex {
			break
		}

		if arr1[arr1InitIndex] == arr2[arr2InitIndex] {
			arr2InitIndex++
		}

		arr1InitIndex++
	}

	if len(arr2) == arr2InitIndex {
		fmt.Println("Subsequence exist")
	} else {
		fmt.Println("No exist")

	}
}
