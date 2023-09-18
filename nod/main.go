package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}
	fmt.Println(euclidGcd(a, b))
}
func euclidGcd(a int, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}
	if a >= b {
		return euclidGcd(a%b, b)
	}

	return euclidGcd(a, b%a)
}
