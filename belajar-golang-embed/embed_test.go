package main_test

import (
	_ "embed"
	"fmt"
	_ "io/fs"
	_ "os"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}