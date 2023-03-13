package main

import (
	"fmt"
	"net-flashbang/controllers"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: net-flashbang <ip-address-or-cidr>")
		os.Exit(1)
	}

	inputValue := os.Args[1]

	if strings.Contains(inputValue, "/") {
		// CIDR-Range eingegeben
		pingResults := controllers.PingRange(inputValue)
		controllers.DisplayPingResults(pingResults)
	} else {
		// IP-Adresse eingegeben
		controllers.Bang()
	}

}
