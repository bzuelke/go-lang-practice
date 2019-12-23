package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Square root of 49 is?", math.Sqrt(49))
	fmt.Println(math.Pi)
}
