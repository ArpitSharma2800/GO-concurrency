package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU units\t", runtime.NumCPU()) //number of CPUs available
	fmt.Println("GO Operating System\t", runtime.GOOS) //operating system for GO environment
	fmt.Println("GO ARCH\t",runtime.GOARCH) // Architecture of GO environment
	fmt.Println("GO ROUTINES\t",runtime.NumGoroutine()) //number of active routines
}
