package main

import (
	"fmt"
)

func main(){
	/*mySlice1 := []int{}
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println(mySlice1)

	mySlice2 := []string{"Go","slices","are","powerful"}
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println(mySlice2)*/

	/*//create a slice from an array
	arr1 :=  [6]int{10,11,12,13,14,15}
	myslice := arr1[2:4]

	fmt.Printf("%v\n",myslice)
	fmt.Printf("%d\n",len(myslice))
	fmt.Printf("%d", cap(myslice))*/

	/*// creathing slices with the make() function
	myslice1 := make([]int,5,10)
	myslice2 := make([]int,5)

	fmt.Printf("myslice1=%v\n",myslice1)
	fmt.Printf("length=%d\n", len(myslice1))
	fmt.Printf("capacity=%d\n",cap(myslice1))
	fmt.Printf("myslice2=%v\n",myslice2)
	fmt.Printf("length=%d\n", len(myslice2))
	fmt.Printf("capacity=%d\n",cap(myslice2))

	// Usage of append()
	myslice3 := append(myslice1,23,27,30,35,40,45,50)
	fmt.Printf("myslice3=%v\n",myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n",cap(myslice3))*/

	numbers:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	//original slice
	fmt.Printf("numbers=%v\n",numbers)
	fmt.Printf("length=%d\n",len(numbers))
	fmt.Printf("capacity=%d\n",cap(numbers))

	//Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy=%v\n", numbersCopy)
	fmt.Printf("length=%d\n", len(numbersCopy))
	fmt.Printf("capacity=%d\n",cap(numbersCopy))



}