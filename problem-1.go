package main

import "fmt"

func find_sum(a, b int) (res int) {
	x := b / a
	res = a * x * (x + 1) / 2
	return
}

func main() {

	fmt.Println(find_sum(3, 999) + find_sum(5, 999) - find_sum(15, 999))

	return
}
