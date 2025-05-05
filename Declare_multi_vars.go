package main

import "fmt"

// If a type keyword is declared ypu can only declare all variables of one type
/*func main(){
	var a,b,c,d int = 1,3,5,7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}*/

// If the type keyword is  not specified, you can declare different types of variables on the same line.
/*func main(){
	var a,b = 6,"Hello"
	var c,d = 7, "World"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}*/

// Go variable declaration in a block
func main(){
	var(
		a int
		b = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// Variable declaration rules

// Camel case : myVariableName 

// Pascal case : MyVariableName

// Snake case : my_variable_name