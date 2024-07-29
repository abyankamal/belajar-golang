package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke ", counter)
	//	counter++
	//}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	fmt.Println("Selesai")

	// perulangan manual
	names := []string{"Eko", "Kurniawan", "Khannedy"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// perulangan otomatis
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// perulangan tanpa index
	for _, name := range names {
		fmt.Println(name)
	}
}