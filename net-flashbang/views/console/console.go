package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// PrintMenu prints the menu to the console
func PrintMenu() {
	fmt.Println(`
	###########################################
	#******* WELCOME TO NET FLASHBANG***********
	#******* CHOOSE YOUR OPTION BELOW *********
	# 1. Ping an IP address
	###########################################
	`)
}

// AskForInput reads from console until line break is available and returns input
func AskForInput() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	response = strings.TrimSpace(response)
	return response
}

// Clear clears the console view
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func AskForIP() string {
	fmt.Println(`
	###########################################
	#******* insert the IP address or DNS*****#
	#******* with/whitout flag ***************#
	###########################################
	`)

	IpAddress := AskForInput()

	return IpAddress
}
