package main

import (
	"fmt"
	"time"
)

func Main() {
	go Spinner(100 * time.Millisecond)
	const n = 10
	fibN := Fib(n)
	fmt.Println(fibN)
}

/*
*

	大写其他模块才能使用
*/
func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func Fib(x int) int {
	if x < 2 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}
