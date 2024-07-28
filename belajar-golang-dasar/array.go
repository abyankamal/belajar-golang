package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Abyan"
	names[2] = "Kamal"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// names[3] = "Apalah" error

	// dynamic array
	
	var values = []int {90, 80, 70, 100, 120}

	fmt.Println(values)

	//ubah data array
	values[2] = 60

    fmt.Println(values)

	// mengetahui panjang array

	var arraylength = len(values)
	fmt.Println("Length of array is", arraylength)
}