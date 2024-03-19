package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
	token := []string{"ETH", "BTC", "USDT", "MATIC"}
	for i := 0; i < len(token); i++ {
		fmt.Println(token[i])
	}

	// to get the index value

	for index, value := range token {
		fmt.Printf("the value of index %v is %v \n", index, value)
	}

}
