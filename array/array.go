package main

import "fmt"

func main() {
	token := [3]string{"ETH", "BTC", "USDT"}
	token[0] = "DEFIPE"
	prices := []int{20, 30, 40}
	prices = append(prices, 20)
	fmt.Println(token)
	fmt.Println(prices)
}
