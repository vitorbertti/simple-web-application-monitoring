package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

	sites := getFileSites()

	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			checkSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}

func checkSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(time.Now())

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "is working normally!")
	} else {
		fmt.Println("Site:", site, "is not working. Status Code:", response.StatusCode)
	}

	fmt.Println("")
}

func getFileSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()

	return sites
}
