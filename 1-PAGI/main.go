package main

import(
	"fmt"
)

var (
	myNum int32
	myNum2 float32
	myNumStr string
	x int8
	y int8
	z int8
)

func main(){
	// Nomer 1
	myNum = 50
	fmt.Printf("%d\n",myNum)

	// Nomer 2
	myNum2 = 51
	fmt.Printf("%f\n",myNum2)

	// Nomer 3
	myNumStr = "50"
	fmt.Printf("%s\n",myNumStr)

	// Nomer 4
	x,y = 10,1
	// Nomer 5
	z= x+y
	fmt.Printf("%d\n",z)
}