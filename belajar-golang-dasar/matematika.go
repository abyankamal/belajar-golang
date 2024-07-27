package main

import "fmt"

func main() {
	var a int8 = 10
	var b int8 = 10
	var c int8 = 5
	var d int8 = 2
	var e int8 = 10

	var result = a + b / c - d * e
	fmt.Println("Hasil:", result)
	
	// assignments operation

	var i = 10
	i++
	fmt.Println("i:", i)

	var j = 10
	j--
	fmt.Println("j:", j)
}