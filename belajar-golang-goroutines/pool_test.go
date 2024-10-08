package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			s := "New"
			return &s
		},
	}

	eko := "Eko"
    kurniawan := "Kurniawan"
    khannedy := "Khannedy"
    pool.Put(&eko)
    pool.Put(&kurniawan)
    pool.Put(&khannedy)


	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*string)
			fmt.Println(*data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}