package main

import "fmt"

func calculator(chn chan int64) {
	a := int64(1)
	b := int64(2)
	for {
		chn <- a
		b += a
		a = b - a
	}
}

func main() {

	sum := int64(0)

	chn := make(chan int64, 1000)

	go calculator(chn)

	for {
		a := <-chn
		if a > 4_000_000 {
			break
		}
		if a%2 == 0 {
			sum += a
		}
	}

	fmt.Println(sum)

	return
}
