package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Eko"
	//person["address"] = "Subang"
	
	person := map[string]string{
		"name":  "John Doe",
        "age":   "30",
        "city":  "New York",
        "hobbies": "reading, playing guitar",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["city"])
	fmt.Println(person["hobbies"])
	fmt.Println(person["food"])

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Eko"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}