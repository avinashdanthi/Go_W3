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

# Go Slices

Slices are similar to arrays, but more powerful and flexible
Like arrays, slices are also used to store multiple values of the same type in a single variable.
However, unlike arrays, the length of a slice can grow and shrink as you see it.
In Go, there are several ways to create a slice:

1.  Using the []datatype{values}format
    2.Create slice from an array
    3.Using the make() function

mySlice=[]int // empty slice of 0 length and o capacity
mySlice=[]int{1,2,3} // slice of integers of length 3 and also the capacity of 3

In Go, there are two functions that can be used to return the length and capacity of a slice:
len() function: returns the length of the slice(the number of elements in the slice)
cap() function: returns the capacity of slice(the number of elements the slice can grow or shrink to)

Create a slice from an array:You create a slice by slicing an array:
var myarray = [length]datatype{values} // an array
mySlice := myarray[start:end] // A slice made from the array

Create a slice with the make() function
slice_name := make([]type,length,capacity)
Note : If the capacity parameter is not defined, it will be equal to length.

Modify slice : Access,Change,Append,Copy
slice_name[index] // Access
slice_name[index] = element
slice_name = append(slice_name, element1, element2,....)
