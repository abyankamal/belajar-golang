package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulus = nilaiAkhir > 80 && absensi > 80

	fmt.Println("Lulus?", lulus)
}