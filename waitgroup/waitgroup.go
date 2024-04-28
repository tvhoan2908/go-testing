package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond)
		fmt.Printf("%d", i)
	}
}

func printLetters2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i < 'A' + 10; i++ {
		time.Sleep(time.Microsecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)

	wg.Wait()
}