package main

import "fmt"

func main() {
	type NoKTP string

	var noKTP1 NoKTP = "1234567890123456"

	var contoh string = "2222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Printf("No KTP: %s\n", noKTP1)
	fmt.Printf("Contoh: %s\n", contoh)
	fmt.Printf("Contoh KTP: %s\n", contohKtp)
}