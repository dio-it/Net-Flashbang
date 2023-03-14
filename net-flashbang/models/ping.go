package models

import (
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

// PingResult ist eine Struktur, die das Ergebnis eines Ping-Befehls enthält.
// IP ist die IP-Adresse des gepingten Hosts.
// Live gibt an, ob der Host auf den Ping reagiert hat.
// Output ist der Textausgabe des Ping-Befehls.
// Error enthält eine Fehlermeldung, wenn beim Ausführen des Ping-Befehls ein Fehler aufgetreten ist.
type PingResult struct {
	IP     string
	Live   bool
	Output string
	Error  error
}

// Ping führt einen Ping-Befehl auf die angegebene IP-Adresse aus und gibt ein PingResult zurück.
func Ping(ip string) PingResult {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	output, err := cmd.Output()

	return PingResult{
		IP:     ip,
		Live:   err == nil && strings.Contains(string(output), "1 received"),
		Output: string(output),
	}
}

// PingRange führt Ping-Befehle auf einem IP-Adressbereich aus und gibt ein Slice von PingResult zurück,
// das die Ergebnisse enthält. Der IP-Adressbereich wird als CIDR-Notation angegeben (z.B. "192.168.0.0/24").
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

	// Startet eine bestimmte Anzahl von Arbeiter-Routinen, um den IP-Bereich zu scannen.
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

	// Fügt jede IP-Adresse in den IP-Bereich in den Kanal für die Arbeiter-Routinen ein.
	for _, ip := range ips {
		workerCh <- ip
	}
	close(workerCh)

	// Wartet, bis alle Arbeiter-Routinen fertig sind, bevor die Ergebnisse zurückgegeben werden.
	wg.Wait()

	return liveIPs
}

// inc erhöht die angegebene IP-Adresse um eins.
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// NewPinger erstellt eine neue Pinger-Instanz, die für die Durchführung von ICMP-Echoanforderungen verwendet werden kann.
// host ist der Hostname oder die IP-Adresse des Ziels.
// count ist die Anzahl der ICMP-Echoanforderungen, die gesendet werden sollen.
// size ist die Größe der Nutzlast in Byte, die mit jeder ICMP-Echoanforderung gesendet werden soll.
// interval ist das Zeitintervall zwischen dem Senden von ICMP-Echoanforderungen.
// timeout ist das Zeitlimit für das Warten auf eine ICMP-Echoantwort.
// ttl ist die "Time to Live"-Feld in der IP-Header der ICMP-Echoanforderung.
// privileged gibt an, ob Pinger mit erhöhten Berechtigungen ausgeführt werden soll, um RAW-Sockets verwenden zu können.

func NewPinger(host string, count int, size int, interval, timeout *time.Duration, ttl *int, privileged bool) (*probing.Pinger, error) {
	// Erstellen einer neuen Pinger-Instanz mit dem angegebenen Host als Ziel
	pinger, err := probing.NewPinger(host)
	if err != nil {
		return nil, err
	}

	// Setzen der Anzahl der ICMP-Echoanforderungen, die gesendet werden sollen
	pinger.Count = count

	// Setzen der Größe der Nutzlast in Byte, die mit jeder ICMP-Echoanforderung gesendet werden soll
	pinger.Size = size

	// Setzen des Zeitintervalls zwischen dem Senden von ICMP-Echoanforderungen
	pinger.Interval = *interval

	// Setzen des Zeitlimits für das Warten auf eine ICMP-Echoantwort
	pinger.Timeout = *timeout

	// Setzen des "Time to Live"-Felds in der IP-Header der ICMP-Echoanforderung
	pinger.TTL = *ttl

	// Setzen der Berechtigungen für die Verwendung von RAW-Sockets
	pinger.SetPrivileged(privileged)

	return pinger, nil

}
