package main

import "fmt"

func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we вone?")
	fmt.Println("Doing some stuff...")
}

func main() {
	foo()
}
