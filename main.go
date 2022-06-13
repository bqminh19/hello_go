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
	fmt.Println("Hello Go")
}
