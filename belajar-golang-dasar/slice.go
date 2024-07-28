package main

import "fmt"

func main() {
	names := [...]string {"Muhammad", "Abyan", "Kamal", "Udin", "Petot"}

	var slice1 []string = names[1:3]
	fmt.Println(slice1)
	slice2 := names[2:4]
	fmt.Println(slice2)
	slice3 := names[1:] // ambil nilai batas atas saja
	fmt.Println(slice3)
	slice4 := names[:2] // ambil nilai batas bawah saja
	fmt.Println(slice4)
	slice5 := names[:] // ambil semua nilai
	fmt.Println(slice5)
}