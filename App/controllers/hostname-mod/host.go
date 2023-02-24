package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s IP-Address\n", os.Args[0])
		os.Exit(1)
	}

	ipAddress := os.Args[1]

	ip := net.ParseIP(ipAddress)
	if ip == nil {
		fmt.Println("Invalid IP address")
		return
	}

	names, err := net.LookupAddr(ip.String())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
