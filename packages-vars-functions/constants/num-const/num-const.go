package main

import "fmt"

const (
	// << is creating a binary number that is 1 followed by 100 zeroes
	Big = 1 << 100
	// Shift it right once more 99 places essentially making it with 2
	Small = Big >> 99
)

func needInt(x int) int { return x + 10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
