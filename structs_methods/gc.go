package main

import (
	"fmt"
	"runtime"
)

func main() {
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)

	fmt.Println("heap after gc. used:", ms.HeapInuse, "free:", ms.GCSys)
}
