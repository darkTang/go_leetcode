package main

import "fmt"

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Println(bagProblem1(weight, value, 4))
}

// 二维dp数组  dp[i][j] i 表示取 0-i 任意物品，j 表示 背包重量
func bagProblem1(weight []int, value []int, bagWeight int) int {
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}
	for i := 0; i <= bagWeight; i++ {
		if weight[0] > i {
			dp[0][i] = 0
		} else {
			dp[0][i] = value[0]
		}
	}

	for i := 1; i < len(weight); i++ {
		for j := 1; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return dp[len(weight)-1][bagWeight]
}

// 一维数组 dp[j] j 表示 背包重量
func bagProblem2(weight []int, value []int, bagWeight int) int {
	return 0
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
