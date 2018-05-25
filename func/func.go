package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("A_i =", i)
		defer func() {
			fmt.Println("B_i =", i)
		}()
		fs[i] = func() {
			fmt.Println("C_i =", i)
		}
	}

	for _, f := range fs {
		f()
	}
}
