package main

import (
	"fmt"
)

func main() {
	//bagProblem5()
	//weight := []int{7, 3, 4, 2}
	//value := []int{3, 6, 9, 1, 2, 4, 10, 1, 3}

	//sort.Slice(value, func(i, j int) bool {
	//	weight[i], weight[j] = weight[j], weight[i]
	//	return value[i] < value[j]
	//})
	//
	//fmt.Println(weight)
	//fmt.Println(value)
	//fmt.Println(quickSort(value))

	v := make([]int, 1, 2)
	fmt.Printf("%p\n", &v)
	fmt.Printf("%p\n", v)
	s := append(v, 1)
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", s)
	v = append(v, 1)
	fmt.Printf("%p\n", &v)
	fmt.Printf("%p\n", v)

}

func quickSort(value []int) []int {
	if len(value) <= 1 {
		return value
	}
	var left []int
	var right []int
	pivot := value[0]
	for i := 1; i < len(value); i++ {
		if value[i] >= pivot {
			right = append(right, value[i])
		} else {
			left = append(left, value[i])
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)

}

func bagProblem5() {
	var m int
	var n int
	var totalValue int
	_, _ = fmt.Scan(&m, &n)
	weight := make([]int, n)
	value := make([]int, n)
	for i := 0; i < n; i++ {
		var w int
		var v int
		_, _ = fmt.Scan(&w, &v)
		weight[i] = w
		value[i] = v
		totalValue += v
	}

	// 贪心算法

	// 价值排序 高 -> 低
	bubbleSort(weight, value)

	// 价值从高往低地来取
	maxValue := 0
	for k := 1; k <= n; k++ {
		total := 0
		totalWeight := m
		totalValue -= value[n-k]
		for i := n - k; i >= 0; i-- {
			if totalWeight >= weight[i] {
				total += value[i]
				totalWeight -= weight[i]
			}
		}
		if total > maxValue {
			maxValue = total
		}

		// 贪心算法优化
		if total >= totalValue {
			break
		}
	}
	fmt.Println(maxValue)
}

func bubbleSort(weight, value []int) {
	for i := 0; i < len(value)-1; i++ {
		for j := 0; j < len(value)-i-1; j++ {
			if value[j] > value[j+1] {
				value[j], value[j+1] = value[j+1], value[j]
				weight[j], weight[j+1] = weight[j+1], weight[j]
			}
		}
	}
}

/*
	10 4
	7  6
	3  3
	4  5
	2  2
*/

// 10
