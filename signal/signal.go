package main

import "fmt"
import "os"
import "os/signal"
import "syscall"
import "sync"
import "time"

func main() {
	var wg sync.WaitGroup
	sigs := make(chan os.Signal, 0)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	go func() {
	FLAG:
		for {
			select {
			case sig := <-sigs:
				fmt.Println(sig)
				done <- true
				wg.Done()
				break FLAG
			case <-time.After(time.Second * 3):
				fmt.Println("Time Out")
				done <- true
				wg.Done()
				break FLAG
			}
		}
	}()

	fmt.Println("Waiting signal")
	fmt.Println("EndFlag:", <-done)
	fmt.Println("exiting")
	wg.Wait()
	time.Sleep(time.Second * 5)
}
