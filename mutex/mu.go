package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func read(i int) {
	fmt.Println("Read:", i)
	wg.Done()
}

func write(i int) {
	fmt.Println("Write:", i)
	wg.Done()
}

func main() {
	wg.Add(4)
	go read(1)
	go read(2)
	go write(3)
	go write(4)
	wg.Wait()
}
