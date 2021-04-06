package main

import "fmt"

func is_prime(a int64) bool {
	for i := int64(2); i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func div_counter(a, b int64) (res int64) {

	for res = 0; a%b == 0; a, res = a/b, res+1 {
	}

	return
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func fast_exp(a, b int64) (res int64) {
	if b == 0 {
		return 1
	}
	res = fast_exp(a, b>>1)
	res *= res
	if b&1 == 1 {
		res *= a
	}
	return
}

func main() {

	res := int64(1)

	for i := int64(2); i <= 20; i++ {
		if is_prime(i) {
			mx := int64(1)
			for j := i; j <= 20; j++ {
				mx = max(mx, div_counter(j, i))
			}
			res *= fast_exp(i, mx)
		}

	}

	fmt.Println(res)

	return
}
