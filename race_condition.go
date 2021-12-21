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
	wg.Add(predator) //creating waitgrp

	for i:=0; i<predator; i++{
		go func() {
			temp := counter
			runtime.Gosched() //it yields the processor, allowing other goroutines to run
			temp++
			counter = temp
			wg.Done()
		}()
		fmt.Println("GO routines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GO routines", runtime.NumGoroutine())
	fmt.Println("GO counter", counter)
}
