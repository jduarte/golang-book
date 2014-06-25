package main

import "fmt"

func main() {
	fmt.Println("Say something: ")
	var something string
	fmt.Scanf("%s", &something)

	fmt.Println("Something is " + something)
}
