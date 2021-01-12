package main

import "fmt"

func main() {
	hashMap := Constructor()
	hashMap.Put(7, 10)
	fmt.Printf("Get 7 = %v\n", hashMap.Get(7))

	hashMap.Put(700007, 20)
	fmt.Printf("Get 7 = %v\n", hashMap.Get(7))

	v1 := hashMap.Get(100)
	fmt.Printf("v1 = %v\n", v1)
	hashMap.Remove(7)
	v1 = hashMap.Get(7)
	fmt.Printf("v1 = %v\n", v1)
}
