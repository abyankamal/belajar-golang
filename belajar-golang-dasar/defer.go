package main

import "fmt"

func logging() {
	fmt.Println("Selesai Melakukan Logging...")
}

func runApplication() {
	defer logging() // dipanggil belakangan dan tetap dieksekusi apapun yang terjadi
	fmt.Println("Mulai Melakukan Aplikasi...")

    // logging() error
}

func main() {
	runApplication()
}