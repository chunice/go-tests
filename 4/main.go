package main

import "fmt"

func josephusRing(n int, m int) int {
	// dp[n] = (dp[n-1]+m) % n
	dp := 0
	for i := 2; i <= n; i++ {
		dp = (dp + m) % i
	}
	return dp
}

func main() {
	ret := josephusRing(3, 2)

	fmt.Printf("ret: %+v", ret)
}
