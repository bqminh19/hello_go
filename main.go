package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHeHe() {
	fmt.Print("Hehe")
}

func sayLoLo() {
	fmt.Print("Lolo")
}

type Point struct {
	x, y int
}

func NewGoRoutine(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			f()
			fmt.Println(" ", i)
			time.Sleep(200 * time.Millisecond)
		}

		wg.Done()
	}()
}
func main() {
	var wg sync.WaitGroup

	defer func() {
		fmt.Println("Before panic")
	}()

	defer wg.Wait()

	NewGoRoutine(&wg, sayLoLo)
	NewGoRoutine(&wg, sayHeHe)

	x := 5
	y := 0
	fmt.Println(x / y)
}
