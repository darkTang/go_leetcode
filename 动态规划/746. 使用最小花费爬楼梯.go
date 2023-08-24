package main

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		dp[i] = isMin(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func isMin(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}
