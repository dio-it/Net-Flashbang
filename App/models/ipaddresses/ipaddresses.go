package ipaddresses

import (
	"flag"
	"fmt"
	"net"
)

func ipAdresse(ipStr string) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		fmt.Printf("%q is not a valid IP address\n", ipStr)
		return
	}
	fmt.Println("IP address:", ip.String())
}

func subnet(s string) {
	_, subnet, err := net.ParseCIDR(s)
	if err != nil {
		fmt.Println("Error parsing subnet:", err)
		return
	}
	fmt.Println("Subnet IP:", subnet.IP.String())
	fmt.Println("Subnet mask:", subnet.Mask.String())
}

func allSub() {
	_, private24, _ := net.ParseCIDR("10.0.0.0/8")
	_, private20, _ := net.ParseCIDR("172.16.0.0/12")
	_, private16, _ := net.ParseCIDR("192.168.0.0/16")

	ips := make([]net.IP, 0)

	for ip := private24.IP; private24.Contains(ip); inc(ip) {
		ips = append(ips, ip)
	}

	for ip := private20.IP; private20.Contains(ip); inc(ip) {
		ips = append(ips, ip)
	}

	for ip := private16.IP; private16.Contains(ip); inc(ip) {
		ips = append(ips, ip)
	}

	fmt.Println("All RFC1918 subnet IP addresses:", ips)
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func ipaddresses() {
	ipStr := flag.String("i", "", "Specify an IP address")
	subnetStr := flag.String("s", "", "Specify a subnet")
	allSubnet := flag.Bool("a", false, "Generate all RFC1918 subnet IP addresses")
	flag.Parse()

	if *ipStr != "" {
		ipAdresse(*ipStr)
	}

	if *subnetStr != "" {
		subnet(*subnetStr)
	}

	if *allSubnet {
		allSub()
	}
}
