package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPU units\t\t", runtime.NumCPU()) //number of CPUs available
	fmt.Println("GO Operating System\t\t", runtime.GOOS) //operating system for GO environment
	fmt.Println("GO ARCH\t\t",runtime.GOARCH) // Architecture of GO environment
	fmt.Println("GO ROUTINES\t\t",runtime.NumGoroutine()) //number of active routines

	wg.Add(1) // adding wait, it will inform system to wait for the task
	go predator()
	zephyrus()

	fmt.Println("task completed successfully")
	fmt.Println("GO ROUTINES\t\t",runtime.NumGoroutine())
	wg.Wait() //this will be called when there is no task left
}

func predator() {
	fmt.Println("Executing predator")
	wg.Done() //this will infirm system that task is completed
}

func zephyrus() {
	fmt.Println("Executing zephyrus")
}
