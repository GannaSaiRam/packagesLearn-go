# [**math**](https://pkg.go.dev/math)
This is the package used to entertain ourself as we are math lover, ain't we? 😉<br>
Like any language, math operations are supported in Go too. In my personal experience we haven't used this math module much, but the learning I did while I was using this package is more than anything. I got very good with syntax, I get to understand defining structs, interfaces, much more. I suggest you to try this module for learning the syntax and show your understanding.<br>
I am sorry I won't be exploring all functions here. It is YOU who have to learn.

### Constants
This math package has already stored few of the constants to make our life easy
- pi
- e
- etc...
### Functions
This math package provides almost all functions we used as child. It has power, absolute, root, ceil, floor, ...., including Inf, NaN <br>
Function | Use
--- | ---
Abs | Get absolute value of given number
Cbrt | Cube Root
Ceil | Ceil of the given float64
Copysign | Convert the given float64 to given sign
Exp | Exponent value
Floor | Floor of the given flaot
Log | Logarithm
Mod | Module of while division between two numbers(sign is equal of dividend)
Pow | Pow is not used for increment but also can used for any kind of powering and rooting
Remainder | Remainder while dividing two float numbers
Round | Rounding the given float64
Sqrt | Square root of given float
Trunc | Returns integer part as float value

# [**math/big**](https://pkg.go.dev/math/big)
This supportes three types of datatypes, float, int, rational. This is used for arthmetic operations. Here we create datatypes and store values as bytes rather than exact values 🤯. I'm not inserting all functions below, please checkout for some more functions in documentations

- ### Float
Function | Use
--- | ---
NewFloat | Initialize the big.Float
Abs | Store given big.Float value and also returns the same fmt.Float
Add | Add two values and stores in self big.Float
Cmp | Comparing with given
Copy | Copying given one to self
Int | Modifies the passed big.Int to float's int part and returns a copy of updated int
MarshalText | Float number to a byte array
Text | There exiting format of priting the float, use one of them
UnmarshalText | Given bytes are updated to big.float
- ### Int
Function | Use
--- | ---
NewInt | Initialize new big.Int
And | And operators between two integers(bitwise)
AndNot | (x)&(!y) operation
Binomial | Its combination value calculator and stores in self 
Cmp | comparing two bit.Ints
Div | Division of ints and store in self
Int64 | Int64 representation of given big.Int
MarshalJSON | Retruns int as byte format
Mul | Multiplication of two big.Ints and store in self
MulRange | 
Neg | change sign of given number
ProbablyPrime | returns if the given number is probably prime or not using Miller-Rabin test
Sign | Sign of given big.Int
Text | retrun in given bease format
UnmarshalJSON | Set the given string to int
- ### Rat
Function | Use
--- | ---
NewRat | Create new instance of rational number
Denom | Denominator of given rational number
FloatString | float value till given precision
Inv | a/b to b/a and stores in self
Num | Numinator of given ration number
Quo | Quotinet when doing division
SetFloat64 | Set given float to self

# [**math/bits**](https://pkg.go.dev/math/bits)
This package is used mostly manipulating bits. Since it directly talk with kernel, so operations are so faster. All functions below are bitwise operations
Function | Use
--- | ---
Add | Adding two numbers catching result of carry also
Div | divison returning quotient and remainder
LeadingZeros | Countingnumber of zeros before the first 1 in binary format of that number of uint
Len | Size of given number in binary format without leading zero.
Mul | Multiplication of number given as arguments
OnesCount | Number of bits with bit value 1
RotateLeft | Rotating given number by given parameter, like using <<, if parameter is negative, the right most will become left most and rotates
# [**math/cmplx**](https://pkg.go.dev/math/cmplx)
The packages is solely for complex number manipulations we do at our school/college.
Function | Use
--- | ---
Abs | Its absolute value of given complex number
Acos, Asin, Acosh, Asinh, ... | Functions supported
Conj | Conjugate of given complex number
Exp | Its nothing but a powering of e to complex number given
NaN | Not-a-Number
Phase | The angle we get from polar notation
Polar | Polar notion of given complex number
Rect | Polar notation to complex number, opposite of Polar
Sqrt |  
# [**math/rand**](https://pkg.go.dev/math/rand)
Random number generator. Please don't use in production level code(use crypto/rand for production).<br>
There are inbuilt functions for generating random number. Get familiar with it.<br>
Seed is used in random number generation to have same sequence of output how many times you run.<br>
- If you are going to run as a server for longer time. This is fine to use.<br>
- But if you are going to start the run each time and you are trying to expect a random number on each run, you won't get the random output, instead you get same output as previous run until you change seed value.

