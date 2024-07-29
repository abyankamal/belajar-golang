package main

import "fmt"

func getFullName() (string, string) {
	return "Abyan", "Kamal"
}

func main(){
	firstname, lastname := getFullName()
	fmt.Println("Hello, my name is", firstname, lastname)

	// menghiraukan return value

	name, _ := getFullName()
	fmt.Println("Hello, my name is", name)
}