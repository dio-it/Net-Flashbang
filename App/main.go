package main

import (
	"flag"
	"fmt"

	"App/models/ipaddresses"
)

func main() {
	ipStr := flag.String("i", "", "Specify an IP address")
	subnetStr := flag.String("s", "", "Specify a subnet")
	allSubnet := flag.Bool("a", false, "Generate all RFC1918 subnet IP addresses")
	flag.Parse()

	// Define the table header
	header := "IP Address"

	// Initialize a list of rows
	rows := [][]string{{header}}

	if *ipStr != "" {
		rows = append(rows, []string{*ipStr})
	}

	if *subnetStr != "" {
		ips, _ := ipaddresses.GetIPsFromSubnet(*subnetStr)
		for _, ip := range ips {
			rows = append(rows, []string{ip.String()})
		}
	}

	if *allSubnet {
		ips, _ := ipaddresses.GetAllRFC1918SubnetIPs()
		for _, ip := range ips {
			rows = append(rows, []string{ip.String()})
		}
	}

	// Calculate the maximum width for each column
	colWidths := make([]int, len(rows[0]))
	for _, row := range rows {
		for i, cell := range row {
			cellWidth := len(cell)
			if cellWidth > colWidths[i] {
				colWidths[i] = cellWidth
			}
		}
	}

	// Print the table
	for _, row := range rows {
		for i, cell := range row {
			fmt.Printf("%-*s ", colWidths[i], cell)
		}
		fmt.Println()
	}
}
