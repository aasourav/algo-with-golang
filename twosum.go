package main

import "fmt"

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	arr1 := []int{-7, -5, -4, 3, 6, 8, 9}
	res := []int{0, 0, 0, 0, 0, 0, 0}

	first := 0
	last := len(arr1) - 1
	resIndex := len(res) - 1

	for {
		if last < first {
			break
		}

		if AbsInt(arr1[first]) < AbsInt(arr1[last]) {
			res[resIndex] = AbsInt(arr1[last]) * AbsInt(arr1[last])
			last--
		} else {
			res[resIndex] = AbsInt(arr1[first]) * AbsInt(arr1[first])
			first++
		}
		resIndex--
	}

	fmt.Println(res)
}
