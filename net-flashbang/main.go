package main

import (
	"fmt"
	"net-flashbang/controllers"
	"os"
	"strings"
)

func main() {
	// Überprüfen, ob die korrekte Anzahl von Argumenten übergeben wurde
	if len(os.Args) != 2 {
		fmt.Println("Usage: net-flashbang <ip-address-or-cidr>")
		os.Exit(1)
	}

	// Speichern des übergebenen Argumentwerts
	inputValue := os.Args[1]

	// Überprüfen, ob der übergebene Wert eine CIDR-Notation ist
	if strings.Contains(inputValue, "/") {
		// CIDR-Range eingegeben
		pingResults := controllers.PingRange(inputValue)
		controllers.DisplayPingResults(pingResults)
	} else {
		// IP-Adresse eingegeben
		controllers.Bang()
	}
}
