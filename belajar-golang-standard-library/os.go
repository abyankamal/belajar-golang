package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // menyimpan data argument
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname() // menyimpan data hostname
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}