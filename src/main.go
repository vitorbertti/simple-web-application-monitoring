package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter your name")
	var name string
	fmt.Scan(&name)

	fmt.Println()

	fmt.Println("Hello", name)

	fmt.Println()

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	var command int
	fmt.Scan(&command)
	fmt.Println("Command was", command)

}
