package controllers

import (
	"flag"
	"fmt"
	"net-flashbang/models"
	"net-flashbang/views/console"
	"os"
	"os/signal"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func PingIP(ip string) models.PingResult {
	return models.Ping(ip)
}

func PingRange(ipRange string) []models.PingResult {
	return models.PingRange(ipRange)
}

func DisplayPingResult(result models.PingResult) {
	console.DisplayPingResult(result)
}

func DisplayPingResults(results []models.PingResult) {
	console.DisplayPingResults(results)
}

func Bang() {
	timeout := flag.Duration("t", time.Second*100000, "")
	interval := flag.Duration("i", time.Second, "")
	count := flag.Int("c", -1, "")
	size := flag.Int("s", 24, "")
	ttl := flag.Int("l", 64, "")
	privileged := flag.Bool("privileged", false, "")
	flag.Usage = func() {
		console.PrintUsage()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		console.PrintUsage()
		return
	}

	host := flag.Arg(0)
	pinger, err := models.NewPinger(host, *count, *size, interval, timeout, ttl, *privileged)
	if err != nil {
		console.PrintError("Failed to create pinger: %v\n", err)
		return
	}

	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *probing.Packet) {
		console.PrintPacketRecv(pkt)
	}
	pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
		console.PrintPacketRecvDuplicate(pkt)
	}
	pinger.OnFinish = func(stats *probing.Statistics) {
		console.PrintPingStats(stats)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())

	if err := pinger.Run(); err != nil {
		console.PrintError("Failed to ping target host: %v\n", err)
	}
}
