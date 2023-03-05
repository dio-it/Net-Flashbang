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

/*func FlagForPing() (*time.Duration, *time.Duration, *int, *int, *int, *bool, string) {
	fmt.Println(`
	Usage:

		ping [-c count] [-i interval] [-t timeout] [--privileged] host

	Examples:

		# ping google continuously
		ping www.google.com

		# ping google 5 times
		ping -c 5 www.google.com

		# ping google 5 times at 500ms intervals
		ping -c 5 -i 500ms www.google.com

		# ping google for 10 seconds
		ping -t 10s www.google.com

		# Send a privileged raw ICMP ping
		sudo ping --privileged www.google.com<

		# Send ICMP messages with a 100-byte payload
		ping -s 100 1.1.1.1
	`)
	Timeout := flag.Duration("t", time.Second*100000, "")
	Interval := flag.Duration("i", time.Second, "")
	Count := flag.Int("c", -1, "")
	Size := flag.Int("s", 24, "")
	Ttl := flag.Int("l", 64, "TTL")
	Privileged := flag.Bool("privileged", false, "")
	flag.Usage = func() {
		fmt.Print(IpAddress)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}
	Host := flag.Arg(0)

	return Timeout, Interval, Count, Size, Ttl, Privileged, Host
}*/
