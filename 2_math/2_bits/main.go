package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(bits.Add(1<<64-1, 2, 1)) // 2 1
	fmt.Println(bits.LeadingZeros(12))   //0(60)1100 => 60 leading zeros
	fmt.Println(bits.Len(8))             //1000 => 4
	fmt.Println(bits.Mul64(2, 3))        // with return 0 6
	fmt.Println(bits.Reverse8(11))       // return 11010000 =
}
