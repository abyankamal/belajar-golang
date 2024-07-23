package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println("Nilai 32-bit:", nilai32)
	fmt.Println("Nilai 64-bit:", nilai64)
	fmt.Println("Nilai 16-bit:", nilai16)

	// ubah string menjadi angka
	var name = "Abyan Kamal"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}