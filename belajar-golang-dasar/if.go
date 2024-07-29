package main

import "fmt"

func main() {
	name := "Eko"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// short statement

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}