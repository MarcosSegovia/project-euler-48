package main

import (
	"fmt"
	"math/big"
)

var result = big.NewFloat(0)

func recursive(a int) *big.Float {

	var bigA *big.Float = big.NewFloat(float64(a))

	if a == 1000 {
		return bigA.SetMantExp(bigA, a)
	}
	result := result.Add(bigA.SetMantExp(bigA, a), recursive(a+1))

	return result
}

func main() {
	number := recursive(1)
	number_to_string := number.Text('f', 0)
	last_ten_digits  := number_to_string[len(number_to_string)-10:]
	fmt.Println(last_ten_digits)
}
