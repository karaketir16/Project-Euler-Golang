package main

import "fmt"

func is_palindrome(a int64) bool {
	digits := []int64{}

	for ; a > 0; a /= 10 {
		digits = append(digits, a%10)
	}

	for ln := len(digits); ln > 1; ln = len(digits) {
		first := digits[0]
		last := digits[ln-1]
		if first != last {
			return false
		}
		digits = digits[:ln-1]
		digits = digits[1:]
	}

	return true
}

func main() {

	largest := int64(0)

	for i := int64(0); i < 1000; i++ {
		for j := int64(0); j < 1000; j++ {
			if a := i * j; is_palindrome(a) && a > largest {
				largest = a
			}
		}
	}

	fmt.Println(largest)

	return
}
