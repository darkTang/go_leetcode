package main

import (
	"fmt"
)

func main() {
	bagProblem7()
}

// 贪心算法实现
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

// 二维dp数组实现 dp[i][j] i 表示取 0-i 任意物品，j 表示 背包重量
func bagProblem6() {
	var m int
	var n int
	_, _ = fmt.Scan(&m, &n)
	weight := make([]int, n)
	value := make([]int, n)
	for i := 0; i < n; i++ {
		var w int
		var v int
		_, _ = fmt.Scan(&w, &v)
		weight[i] = w
		value[i] = v
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m+1)
	}

	// 初始化dp数组
	for i := 0; i <= m; i++ {
		if i >= weight[0] {
			dp[0][i] = value[0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= m; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = maxNum2(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println("最大价值：")
	fmt.Println(dp[n-1][m])
	fmt.Println("dp数组：")
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			fmt.Printf("%d  ", dp[i][j])
		}
		fmt.Println()
	}
}

// 一维dp数组实现 dp[j] j 表示 背包重量
func bagProblem7() {
	var m int
	var n int
	_, _ = fmt.Scan(&m, &n)
	weight := make([]int, n)
	value := make([]int, n)
	for i := 0; i < n; i++ {
		var w int
		var v int
		_, _ = fmt.Scan(&w, &v)
		weight[i] = w
		value[i] = v
	}

	dp := make([]int, m+1)
	for i := 0; i < n; i++ {
		for j := m; j >= 0; j-- {
			if j >= weight[i] {
				dp[j] = maxNum2(dp[j], dp[j-weight[i]]+value[i])
			}
		}
		// 输出每次的dp数组
		for k := 0; k <= m; k++ {
			fmt.Printf("%d  ", dp[k])
		}
		fmt.Println()
	}
	fmt.Println(dp[m])

}

func maxNum2(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

/*
   【输入】
		10 4
		7  6
		3  3
		4  5
		2  2
    【输出】
		10
*/
/*
   【输入】
		10 4
		2  1
		3  3
		4  5
		7  9
    【输出】
		12
*/
