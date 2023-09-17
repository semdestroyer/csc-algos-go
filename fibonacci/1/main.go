package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	fmt.Println(fibbonaci(n))
}

func fibbonaci(n int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	if n <= 2 {
		return arr[n-1]
	}
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n-1]
}
