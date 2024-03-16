package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func HeavyFunction(wg *sync.WaitGroup) {
	defer wg.Done()
}

func HeavyFunction2(wg *sync.WaitGroup) {
	defer wg.Done()
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	// fmt.Println("In main()")
	// go longWait()
	// go shortWait()
	// fmt.Println("about to sleep in main()")
	// time.Sleep(10 * 1e9)
	// fmt.Println("end of main()")
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go HeavyFunction(wg)
	go HeavyFunction2(wg)
	wg.Wait()
	fmt.Println("all finished")
}

func longWait() {
	fmt.Println("Beggining longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("End longWait()")
}

func shortWait() {
	fmt.Println("Beggining shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("End shortWait()")
}
