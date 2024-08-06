package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := SayHello("World")
    if result!= "Hello World" {
        t.Error("Expected 'Hello World', got", result)
    }
}

func TestHelloEko(t *testing.T) {
	result := SayHello("Eko")
    if result!= "Hello Eko" {
        t.Error("Expected 'Hello Eko', got", result)
    }
}