package main

import "fmt"

type fn func(int)
type ck func(int) bool

func filter(p int, c ck) ck {
	return func(n int) bool {
		if c == nil {
			return true
		}
		if n%p == 0 {
			return false
		}
		return c(n)
	}
}

func prime(n, max int, c ck, f fn) func(int, int, ck) {
	if n > max {
		return nil
	}
	if n == 2 || c(n) {
		f(n)
		return prime(n+1, max, filter(n, c), f)
	}
	return prime(n+1, max, c, f)
}

func Prime(max int, f fn) {
	prime(2, max, nil, f)
}

func main() {
	Prime(100, func(n int) {
		fmt.Println(n)
	})
}
