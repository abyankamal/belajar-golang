package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	// cara iterasi
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// cara manual

	//data.Value = "Value 1"
	//
	//data = data.Next()
	//data.Value = "Value 2"
	//
	//data = data.Next()
	//data.Value = "Value 3"
	//
	//data = data.Next()
	//data.Value = "Value 4"
	//
	//data = data.Next()
	//data.Value = "Value 5"

	// memanggil fungsi dari ring
	data.Do(func(value any) {
		fmt.Println(value)
	})
}