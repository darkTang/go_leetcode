package main

import "fmt"

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Println(bagProblem(weight, value, 4))
}

func bagProblem(weight []int, value []int, bagWeight int) int {
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

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
