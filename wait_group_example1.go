package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("CPU units\t\t", runtime.NumCPU()) //number of CPUs available
	fmt.Println("GO ROUTINES\t\t",runtime.NumGoroutine()) //number of active routines

	wg.Add(3)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("the value of i = ", i)
		}
		wg.Done()
	}()

	go func() {
		fmt.Println("Zyphyrus will be beast")
		wg.Done()
	}()

	go func() {
		k := 10
		if k != 0 {
			fmt.Println("Be the zero that increase value X10")
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all tasks completed successfully")

}
