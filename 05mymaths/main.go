package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var my1 int = 2
	// var my2 float64 = 4.5
	// fmt.Println("The sum is: ", my1+int(my2))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto

	myR, _ := rand.Int(rand.Reader, big.NewInt(10))

	fmt.Println(myR)

}
