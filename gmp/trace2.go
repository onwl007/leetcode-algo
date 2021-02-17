package main

import (
	"fmt"
	"time"
)

// GODEBUG=schedtrace=1000 ./trace2
func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello GMP")
	}
}
