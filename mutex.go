package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("GO CPUs", runtime.NumCPU())
	fmt.Println("GO routines", runtime.NumGoroutine())
	counter := 0
	const predator = 100
	var wg sync.WaitGroup
	var mg sync.Mutex //mutex package

	wg.Add(predator)

	for i:=0; i<predator; i++{
		go func() {
			mg.Lock() //locking so that it will block other routines to access while running
			temp := counter
			temp++
			counter = temp
			mg.Unlock() //it will unlock for other routines to access
			wg.Done()
		}()
		fmt.Println("GO routines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GO routines", runtime.NumGoroutine())
	fmt.Println("GO counter", counter)
}

