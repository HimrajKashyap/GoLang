package main

import "fmt"

func main() {
	fmt.Println("Hello" + "World")
	count := 10
	lastname := "Smith"
	const MAX = 10000

	fmt.Println(MAX) //Constsnts are immutable values that donot change during execution
	fmt.Println(count)
	fmt.Println(lastname)
}
