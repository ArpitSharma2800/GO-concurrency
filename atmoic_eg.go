package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 //creating atomic of int64
	const predator = 100
	var wg sync.WaitGroup
	wg.Add(predator) //creating waitgrp

	for i:=0; i<predator; i++{
		go func() {
			atomic.AddInt64(&counter,1) //adding pointer to the counter with value 1
			fmt.Println("Counter", atomic.LoadInt64(&counter)) //loading pointer to the counter
			wg.Done()
		}()
		fmt.Println("GO routines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GO counter", counter)
}

