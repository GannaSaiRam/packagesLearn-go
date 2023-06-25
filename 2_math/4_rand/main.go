package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("Output number lies in range [0, %v): %v\n", 100, rand.Int63n(100))
	fmt.Printf("Output number lies in range [0, %v): %v\n", 100, rand.Int31n(100))

	r := rand.New(rand.NewSource(1))
	fmt.Println(r.Float32())
}
