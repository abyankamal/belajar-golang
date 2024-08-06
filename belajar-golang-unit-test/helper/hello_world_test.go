package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHelloAssertion(t *testing.T) {
    result := SayHello("World")
    assert.Equal(t, "Hello World", result, "They Should Be Equal")
    
    fmt.Println("Test SayHello Assertion Done")
}

func TestHelloWorld(t *testing.T) {
	result := SayHello("World")
    if result!= "Hello World" {
        t.Error("Expected 'Hello World', got", result)
    }

    fmt.Println("Test Hello World Done")
}

func TestHelloEko(t *testing.T) {
	result := SayHello("Eko")
    if result!= "Hello Eko" {
        t.Fatal("Expected 'Hello Eko', got", result)
    }

    fmt.Println("Test Hello World Done")
}