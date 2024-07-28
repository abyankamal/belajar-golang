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

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// make slice

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"
	// newSlice[2] = "Eko" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy slice

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}