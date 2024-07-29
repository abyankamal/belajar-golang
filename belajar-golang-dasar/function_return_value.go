package main

import "fmt"

func getHello(name string) string {
	hello := fmt.Sprintf("Hello, %s!", name)
	return hello
}

func main(){
	result := getHello("Abyan")
	fmt.Println("Result: ", result)
}