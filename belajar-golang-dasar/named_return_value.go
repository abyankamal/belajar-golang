package main

import "fmt"

func getCompletedName() (firstname, middlename, lastname string) {
	firstname = "Muhammad"
	middlename = "Abyan"
	lastname = "Kamal"

	return firstname, middlename, lastname
}

func main() {
	a, b, c := getCompletedName()
	fmt.Printf("Hello, %s %s %s!\n", a, b, c)
}