package main

import "fmt"

func fibo_fact(n int64) (num int64) {
	if n <= 2 {
		num = 1
	} else {
		num = fibo_fact(n-1) + fibo_fact(n-2)
	}
	return
}

func main() {
	var en int64 = 50
	en = fibo_fact(en)
	fmt.Println(en)
}
