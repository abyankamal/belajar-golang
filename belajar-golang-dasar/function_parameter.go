package main

import "fmt"

func sayHelloTo(firstname string, secondname string){
	fmt.Printf("Hello, %s %s!\n", firstname, secondname)
    fmt.Println("Hello,", firstname, secondname, "!")
}
func main() {
	sayHelloTo("Abyan", "Kamal")
}