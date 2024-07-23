package main

import "fmt"

func main() {
	// deklarasi tunggal
	// const firstName = "Abyan"
	// const lastName = "Kamal"

	// deklarasi banyak const 
	const (
		firstName = "Abyan"
        lastName  = "Kamal"
	)

	// ubah isi variabel
	// firstName = "Abyan Umar" error

    // tampilkan isi variabel
    fmt.Println("First Name:", firstName)
    fmt.Println("Last Name:", lastName)
}