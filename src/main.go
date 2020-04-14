package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoringTimes = 3
const delay = 5

func main() {
	fmt.Println("Welcome")
	fmt.Println("")

	for {
		showMenu()
		command := setCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs")
		case 0:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	fmt.Println("")
}

func setCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := []string{"https://www.google.com.br"}

	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			checkSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func checkSite(site string) {
	response, _ := http.Get(site)
	fmt.Println(time.Now())

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "is working normally!")
	} else {
		fmt.Println("Site:", site, "is not working. Status Code:", response.StatusCode)
	}

	fmt.Println("")
}
