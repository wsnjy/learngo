package belajargolanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("totalCPU", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("totalThread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine)
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("totalCPU", totalCPU)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("totalThread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine)
}
