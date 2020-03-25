package main

func maxProfit(prices []int) int {
	// 三维数组  n: 天， k: 买了几笔 , s: 持有状态  1持有  0不持有
	// dp[n][k][0] = max(dp[n-1][k][0], dp[n-1][k][1] + value )
	// dp[n][k][1] = max(dp[n-1][k][1], dp[n-1][k-1][0] - value )

	if len(prices) < 2 {
		return 0
	}
	dp := make([][][]int, len(prices))
	dealNum := 2
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, dealNum+1)
		for j := 0; j < dealNum+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][1] = -prices[0]
	dp[0][1][1] = -prices[0]
	dp[0][2][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= dealNum; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][2][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
