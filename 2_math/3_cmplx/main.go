package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println(cmplx.Abs(1 + 2i))
	fmt.Println(cmplx.Conj(4 + 7.9i))
	fmt.Println(cmplx.Sqrt(-1) == 1i)
}
