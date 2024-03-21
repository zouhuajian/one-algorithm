package arr

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
// 121. 买卖股票的最佳时机
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
func maxProfit(prices []int) int {
	profit := 0
	minPrice := math.MaxInt32

	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			if prices[i]-minPrice > profit {
				profit = prices[i] - minPrice
			}
		}
	}
	return profit
}

// 使用动态规划实现
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/solutions/38477/bao-li-mei-ju-dong-tai-gui-hua-chai-fen-si-xiang-b/
// https://blog.csdn.net/lw_power/article/details/103772951
func maxProfitDP(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	dp := make([][2]int, l)
	// dp[i][0] 下标为 i 这天结束的时候，不持股，手上拥有的现金数
	// dp[i][1] 下标为 i 这天结束的时候，持股，手上拥有的现金数
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < l; i++ {
		price := prices[i]
		dp[i][0] = getMax(dp[i-1][0], dp[i-1][1]+price)
		dp[i][1] = getMax(dp[i-1][1], -price)
	}
	return dp[l-1][0]
}

func maxProfitDP2(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	cash := 0
	hold := -prices[0]
	for i := 1; i < l; i++ {
		price := prices[i]
		cash = getMax(cash, hold+price)
		hold = getMax(hold, -price)
	}
	return cash
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
