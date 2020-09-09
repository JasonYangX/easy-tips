package main

import "fmt"

// bigTopHeap 大顶堆
func bigTopHeap(input []int) []int {
	aroundFind := false
	lenInput := len(input)
	i := 0
	for {
		if 2*(i+1) > lenInput {
			if !aroundFind {
				return input
			}
			i = 0
			aroundFind = false
			continue
		}
		if input[i] < input[2*i+1] {
			tmp := input[i]
			input[i] = input[2*i+1]
			input[2*i+1] = tmp
			aroundFind = true
		}
		if 2*(i+1) < lenInput && input[i] < input[2*i+2] {
			tmp := input[i]
			input[i] = input[2*i+2]
			input[2*i+2] = tmp
			aroundFind = true
		}
		i++
	}
	return input
}

// lessTopHeap 小顶堆
func lessTopHeap(input []int) []int {
	aroundFind := false
	lenInput := len(input)
	i := 0
	for {
		if 2*(i+1) > lenInput {
			if !aroundFind {
				return input
			}
			i = 0
			aroundFind = false
			continue
		}
		if input[i] > input[2*i+1] {
			tmp := input[i]
			input[i] = input[2*i+1]
			input[2*i+1] = tmp
			aroundFind = true
		}
		if 2*(i+1) < lenInput && input[i] > input[2*i+2] {
			tmp := input[i]
			input[i] = input[2*i+2]
			input[2*i+2] = tmp
			aroundFind = true
		}
		i++
	}
	return input
}

func main() {
	input := []int{2, 7, 26, 25, 19, 17, 1, 90, 3, 36}
	fmt.Println(bigTopHeap(input))
	fmt.Println(lessTopHeap(input))
}
