package main

import "fmt"

type fmtPrintingStruct struct {
	X int
	Y float64
	Z string
	W struct {
		A int
		B complex128
	}
}

func (f fmtPrintingStruct) String() string {
	return fmt.Sprintf(
		"X: %v\nY: %v\nZ: %v\nW.A: %v\nW.B: %v\n",
		f.X, f.Y, f.Z, f.W.A, f.W.B,
	)
}

func main() {
	fmt.Print("Hello, ") // Understand the difference between Print and Println
	fmt.Println("I like learning "+"Go", ". I like Go because of "+"concurrency")
	// For each comma separated variables, space will be inserted
	fmt.Printf("%d\n", 10) // Number of verbs, we define inside should be equal to number of arguments

	// Ordering
	fmt.Printf("I am %[2]v years old. I born in %[3]v. I have %[1]v years work experience\n", 7, 27, 1995)
	fmt.Printf("%[3]*.[2]*[1]f\n", 10.5123, 4, 6) // Is identical like %6.4f
	fmt.Printf("%[1]v %#[1]v\n", 10)

	var err error = fmt.Errorf("this is custom error")
	fmt.Println(err)

	var st fmtPrintingStruct
	fmt.Printf("%[1]v, %+[1]v\n", st) // The difference you can see clearly in the outputs

	var urlFormat = "https://%s/%s"
	actualUrl := fmt.Sprintf(urlFormat, "go.dev", "play")
	fmt.Println(actualUrl)

	// Stringer implementation
	st = fmtPrintingStruct{X: 7, Y: 3.14, Z: "Pi", W: struct {
		A int
		B complex128
	}{A: 3, B: 5 + 7i}}
	fmt.Printf("%v", st)
}
