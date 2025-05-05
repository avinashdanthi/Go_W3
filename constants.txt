package main

import "fmt"

/*const PI float32 = 3.14 // typed constant

const G = 9.8 // untyped constant

func main(){
	fmt.Println("PI value is", PI)
	fmt.Println("Gravity is",G)
}*/

// Multiple constants declaration

func main(){
	const(
		A int = 1
		B = 3.14
		C = "Hi!"
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}