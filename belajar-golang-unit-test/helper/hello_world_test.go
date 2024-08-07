package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Muhammad", func(t *testing.T) {
		result := SayHello("Muhammad")
		require.Equal(t, "Hello Muhammad", result, "Result must be 'Hello Muhammad'")
	})
	t.Run("Abyan", func(t *testing.T) {
		result := SayHello("Abyan")
		require.Equal(t, "Hello Abyan", result, "Result must be 'Hello Abyan'")
	})
	t.Run("Kamal", func(t *testing.T) {
		result := SayHello("Kamal")
		require.Equal(t, "Hello Kamal", result, "Result must be 'Hello Kamal'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
    if runtime.GOOS == "darwin" {
        t.Skip("This test is only for Windows")
    }

    result := SayHello("World")
    require.Equal(t, "Hello World", result, "They Should Be Equal")
}

func TestSayHelloAssertion(t *testing.T) {
    result := SayHello("World")
    assert.Equal(t, "Hello World", result, "They Should Be Equal")
    
    fmt.Println("Test SayHello Assertion Done")
}

func TestSayHelloRequire(t *testing.T) {
    result := SayHello("Apaan")
    require.Equal(t, "Hello World", result, "They Should Be Equal")
    
    fmt.Println("Test SayHello Not Printed")
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