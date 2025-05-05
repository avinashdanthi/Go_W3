# What is Go ?

a. Go is cross-platform, open source programming language
b. Go can be used to create high performance applications
c. Go is a fast, statically typed, compiled language known for its simplicity and efficiency.
d. Go was developed at Google by three members Robert,Rob and Kim in 2007
e. Go syntax is similar to C++

# What is Go used for ?

a. Web development (server-side)
b. Developing network-based programs
c. Developing cross-platform enterprise applications
d. cloud-native development

# Why use Go ?

a. Go has fast run time and compilation time
b. Go supports concurrency
c. Go has memory management (Go has garbage collection and automatic memory management, so you don’t have to manually free memory.)
d. Go works on different platforms(Windows, Mac, lINUX, Raspberry Pi,etc.)

# Go syntax:

A go file consists of the following parts.
a. Package declaration
b. Import pacakages
c. Functions
d. Statements and expressions

# Go Variable Types

int, float_thirty_two, string, bool
In Go there are two ways to declare a variable:
1.with var keyword var variablename type = value 2. With the := sign variablename := value

# Go Constants

The const keyword declares the variable as "constant", which means that it is unchangeable and read only.
const CONSTNAME type = value
Constants follow the same naming convention as variables
Constants are of two types "typed" eg: const A int = 1 amd "untyped" eg: const A = 1

# Go Output functions

1. Print("Prints the raw format") 2. Println("Same as prinf but adds a new line at the end") 3.Printf("It formats its argument based on the given formatting verb")
   General Formatting verbs:
   %v - Prints the value in default format
   %#v - Prints the value in Go syntax format
   %T - Prints the type of the value
   %% - Prints the % sign

# Go Arrays

Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
In Go there are two types of array declaration:
1.With the var keyword
var array_name = [length]datatype{values} // here length is defined
or
var array_name = [...]datatype{values} // here length is inferred 2. With the := sign
array_name := [length]datarype{values} // here the length is defined
or
array_name := [...]datatype{values} // here length is inferred
