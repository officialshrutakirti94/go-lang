package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	message := "20 eth credited"
	fmt.Println(strings.Contains(message, "ETH"))
	fmt.Println(strings.ReplaceAll(message, "20", "100"))
	fmt.Println(strings.ToUpper(message))
	fmt.Println(strings.Index(message, "th"))
	fmt.Println(strings.Split(message, ""))

	price := []int{40, 30, 24, 89, 20, 10, 5}
	sort.Ints(price)
	fmt.Println(price)
	index := sort.SearchInts(price, 40)
	fmt.Println("40 found at:", index)

	token := []string{"eth", "btc", "matic", "weth", "wbtc", "usdt"}
	sort.Strings(token)
	fmt.Println(token)
	fmt.Println(sort.SearchStrings(token, "matic"))
}
