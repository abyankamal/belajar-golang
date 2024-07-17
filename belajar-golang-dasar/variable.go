package main

import "fmt"

func main() {
	// Deklarasi variabel
	// var name = "Abyan Kamal"
	// fmt.Println("Hello, my name is", name)

	// name = "Abyan Akmal"
	// fmt.Println("Hello, my name is", name)

	// tanpa deklarasi variabel
	name := "Abyan Kamal"
	fmt.Println("Hello, my name is", name)

	name = "Abyan Akmal"
	fmt.Println("Hello, my name is", name)

	// bikin banyak variabel
	var (
        firstName = "Abyan"
        lastName  = "Kamal"
    )
    fmt.Println("Hello, my name is", firstName, lastName)
}