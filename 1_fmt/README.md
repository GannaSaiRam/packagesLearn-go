# [**fmt**](https://pkg.go.dev/fmt)
This is the basic package to run this language. From the start of printing on console to write logs into a file, this package can be used. To make it simple, this is a printer package(for beginners).

## **Basic Format Printing**
Verb/Flag | DataType
--- | ---
%d | Any integer(uint, int32, int64, int)
%b | Print given number in binary format
%x/%X | Print integer in hexadecimal format
%o/%O | Print number in octal format
%s | String
%t | Boolean(true / false)
%T | Type of data in string format
%f / %F | Float number
%g | Any complex number, float32
**%v** | Any datatype (we won't require anything above, this is enough)
%p | chan, pointer

_If you like engineering rather being doctor, **%v** is the one for you._😋

### **Format printing with little tweaking**
I would like to explain this with **%v**.<br>
<strong><i>+</i></strong>:<br>
If we want to print any struct with struct.<br>
<strong><i>Printing Order</i></strong>:<br>
Sometimes we don't need to have same order of parameters of verb. And use the same value to print multiple times, we use this idea.
- ("%[1]v, %+[1]v", var) &equiv; ("%v, %+v", var, var)
- ("%[3]v > %[2]v > %[1]v", 5, 10, 54) &equiv; 54 > 10 > 5

## **Use case**
Printing is genereally used for check the status and see the expected result or not. But most of the use case is string operation (in development).
### **Strings**
Till now only printing we have talked about which returns nothing. But generally in real development, we actually require string output.<br>
We do that by using _fmt.Sprintf, fmt.Sprint, fmt.Sprintln_
### **Error**
We can format an error using this fmt.

#### *More Information*
If we want to fmt the way of printing a string we want. We need to define a method String i.e implementing <u>Stringer</u> interface.