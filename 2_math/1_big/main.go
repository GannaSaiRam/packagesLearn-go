package main

import (
	"fmt"
	"math/big"
)

func main() {

	x := big.NewFloat(3.9)
	// Operations
	fmt.Printf("Converting -6.8 to abs: %v\n", x.Abs(new(big.Float).SetFloat64(-6.8)))

	y := new(big.Int)
	y.Add(big.NewInt(3), big.NewInt(4))
	fmt.Println(y.String(), y)

	z := big.NewRat(4, 5)
	fmt.Println(z)

	mantessa := big.NewFloat(3)
	mantessa.SetMantExp(mantessa, mantessa.MantExp(nil)+3) // 1.5*2**3
	fmt.Println(mantessa)

	// Float
	floatBig1 := big.NewFloat(7.5)
	fmt.Println(floatBig1.Abs(big.NewFloat(-6.7)))

	floatBig2 := big.NewFloat(4.5)
	inti := big.NewInt(3)
	d, a := floatBig2.Int(inti)
	fmt.Println(&inti, floatBig2, &d, a)

	floatBig3 := big.NewFloat(1.05)
	textOut, err := floatBig3.MarshalText()
	fmt.Println(floatBig3, textOut, err)

	err = floatBig3.UnmarshalText([]byte("5.903"))
	fmt.Println(floatBig3, floatBig3.Text('e', 4), err)

	// Int
	intBig1 := big.NewInt(11)
	fmt.Println(intBig1, intBig1.Bit(2), intBig1.BitLen(), intBig1.Bits(), intBig1.Bytes())
	fmt.Println(intBig1.MarshalText())
	fmt.Println(intBig1.MarshalJSON())

	// Rat
	ratBig1 := big.NewRat(7, 5)
	fmt.Println(ratBig1.FloatString(5), ratBig1.Quo(big.NewRat(10, 3), big.NewRat(7, 5)))
}
