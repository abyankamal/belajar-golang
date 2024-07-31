package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover() // menangkap pesan panic
	fmt.Println("terjadi panic", message) 
}

func runApp(error bool) {
	defer endApp()
	if error {
        panic("Terdapat Error") // dijalankan ketika error
    }
}

func main() {
	// runApp(false) bakalan jalan
	runApp(true)
	fmt.Println("Eko Kurniawan Khannedy")
}