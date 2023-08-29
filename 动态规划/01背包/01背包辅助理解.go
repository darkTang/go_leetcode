/*
M公斤背包，n件物品，重量分别为W1,W2,...,Wn，价值分别为C1,C2,...,Cn，求获得的最大价值。
【输入】

	第一行：两个数，M背包容量和N物品数量
	第2～N+1行：每行两个整数，表示每个物品的重量和价值

【输出】

	仅一行，一个数，表示最大价值

【输入样例】

	10 4
	2  1
	3  3
	4  5
	7  9

【输出样例】

	12
*/
package main

import "fmt"

func main() {
	bagProblem4()
}

func bagProblem3() {
	var m int
	var n int
	_, _ = fmt.Scanf("%d %d", &m, &n)
	weight := make([]int, n+1) // 为了更好的初始化
	value := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var w, v int
		_, _ = fmt.Scanf("%d %d", &w, &v)
		weight[i] = w
		value[i] = v
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = maxNum(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println("最大价值：")
	fmt.Println(dp[n][m])
	fmt.Println("dp数组：")
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			fmt.Printf("%d  ", dp[i][j])
		}
		fmt.Println()
	}
}

// 一维dp数组，滚动数组
func bagProblem4() {
	var m int
	var n int
	_, _ = fmt.Scanf("%d %d", &m, &n)
	weight := make([]int, n+1) // 为了更好的初始化
	value := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var w, v int
		_, _ = fmt.Scanf("%d %d", &w, &v)
		weight[i] = w
		value[i] = v
	}
	dp := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := m; j >= 1; j-- {
			if j >= weight[i] {
				dp[j] = maxNum(dp[j], dp[j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println("最大价值：")
	fmt.Println(dp[m])
}

func maxNum(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
