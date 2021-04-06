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

func divisors(chn chan int64, a int64) {
	for i := int64(2); i*i <= a; i++ {
		if a%i == 0 {
			chn <- i
			chn <- a / i
		}
	}
	close(chn)
}

func main() {

	chn := make(chan int64, 1000)

	go divisors(chn, 600_851_475_143)

	largest := int64(0)

	for a := range chn {
		if is_prime(a) {
			largest = a
		}
	}

	fmt.Println(largest)

	return
}
