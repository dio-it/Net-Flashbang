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

	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)

	}
	fmt.Println("----------------------interfaces-----------------------")
	fmt.Println(interfaces)
	fmt.Println("----------------------interfaces-----------------------")

	var macAddresses []net.HardwareAddr
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)

		}
		fmt.Println("----------------------interface-----------------------")
		fmt.Println(iface)
		fmt.Println("----------------------interface-----------------------")
		fmt.Println("----------------------addrs-----------------------")
		fmt.Println(addrs)
		fmt.Println("----------------------addrs-----------------------")

		fmt.Printf("Interface: %s\n", iface.Name)
		for _, addr := range addrs {
			fmt.Printf("- IP address: %s\n", addr.String())
			if ipnet, ok := addr.(*net.IPNet); ok {
				if ipnet.IP.Equal(net.ParseIP(ipAddress)) && !ipnet.IP.IsLoopback() {
					// This interface has the IP address we are looking for
					macAddresses = append(macAddresses, iface.HardwareAddr)
				}
			}
		}
	}

	if len(macAddresses) > 0 {
		fmt.Printf("MAC address(es) for %s: %v\n", ipAddress, macAddresses)
	} else {
		fmt.Fprintf(os.Stderr, "Error: no interface found for IP address %s\n", ipAddress)
		os.Exit(1)
	}
}
