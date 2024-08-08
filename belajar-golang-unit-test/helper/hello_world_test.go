package helper

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000) // Generate random integers between 0 and 999
	}
	return slice
}

func BenchmarkBubbleSort(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				slice := generateRandomSlice(size)
				b.StartTimer()

				BubbleSort(slice)
			}
		})
	}
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		},
		{
			name:    "Kurniawan",
			request: "Kurniawan",
		},
		{
			name:    "EkoKurniawanKhannedy",
			request: "Eko Kurniawan Khannedy",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Eko")
		}
	})
	b.Run("Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Kurniawan")
		}
	})
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Eko")
	}
}

func BenchmarkSayHelloKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Kurniawan")
	}
}

func TestTableSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Khannedy",
			request:  "Khannedy",
			expected: "Hello Khannedy",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

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