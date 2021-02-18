package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	catch := make(chan int, 1)
	fishch := make(chan int, 1)
	dogch := make(chan int, 1)
	fishch <- 1

	go cat(fishch, catch)
	go dog(catch, dogch)
	go fish(dogch, fishch)

	wg.Add(3)
	wg.Wait()
}

func cat(fishch, catch chan int) {
	defer wg.Done()
	defer close(catch)
	for i := 0; i < 100; i++ {
		<-fishch
		fmt.Println("cat")
		catch <- 1
	}
}

func dog(catch, dogch chan int) {
	defer wg.Done()
	defer close(dogch)
	for i := 0; i < 100; i++ {
		<-catch
		fmt.Println("dog")
		dogch <- 1
	}
}

func fish(dogch, fishch chan int) {
	defer wg.Done()
	defer close(fishch)
	for i := 0; i < 100; i++ {
		<-dogch
		fmt.Println("fish")
		fishch <- 1
	}
}
