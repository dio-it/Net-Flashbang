package models

import (
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

type PingResult struct {
	IP     string
	Live   bool
	Output string
	Error  error
}

func Ping(ip string) PingResult {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	output, err := cmd.Output()

	return PingResult{
		IP:     ip,
		Live:   err == nil && strings.Contains(string(output), "1 received"),
		Output: string(output),
	}

}

func PingRange(ipRange string) []PingResult {
	ip, ipnet, err := net.ParseCIDR(ipRange)
	if err != nil {
		return []PingResult{}
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var liveIPs []PingResult
	numWorkers := 1000
	workerCh := make(chan string, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ip := range workerCh {
				result := Ping(ip)
				if result.Live {
					mu.Lock()
					liveIPs = append(liveIPs, result)
					mu.Unlock()
				}
			}
		}()
	}

	for _, ip := range ips {
		workerCh <- ip
	}
	close(workerCh)

	wg.Wait()

	return liveIPs
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func NewPinger(host string, count int, size int, interval, timeout *time.Duration, ttl *int, privileged bool) (*probing.Pinger, error) {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		return nil, err
	}

	pinger.Count = count
	pinger.Size = size
	pinger.Interval = *interval
	pinger.Timeout = *timeout
	pinger.TTL = *ttl
	pinger.SetPrivileged(privileged)

	return pinger, nil
}
