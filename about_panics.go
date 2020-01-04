package go_koans

import "fmt"

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic?")
		}
	}()

	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
