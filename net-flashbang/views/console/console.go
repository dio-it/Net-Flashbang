package console

import (
	"fmt"
	"os"

	"net-flashbang/models"

	probing "github.com/prometheus-community/pro-bing"
)

func DisplayPingResult(result models.PingResult) {
	if result.Live {
		fmt.Printf("IP %s is live.\n", result.IP)
	} else {
		fmt.Printf("IP %s is not responding.\n", result.IP)
	}
	fmt.Println(result.Output)
}

func DisplayPingResults(results []models.PingResult) {
	fmt.Println("Live IPs:")
	for _, result := range results {
		DisplayPingResult(result)
	}
}

var usage = `
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
    sudo ping --privileged www.google.com

    # Send ICMP messages with a 100-byte payload
    ping -s 100 1.1.1.1
`

func PrintUsage() {
	fmt.Print(usage)
}

func PrintPacketRecv(pkt *probing.Packet) {
	fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
		pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
}

func PrintPacketRecvDuplicate(pkt *probing.Packet) {
	fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
		pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
}

func PrintPingStats(stats *probing.Statistics) {
	fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
	fmt.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n",
		stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
	fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
		stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
}

func PrintError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}
