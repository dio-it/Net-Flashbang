package controllers

import (
	"flag" // Paket für Kommandozeilenargumente
	"fmt"
	"net-flashbang/models"
	"net-flashbang/views/console"
	"os"
	"os/signal"
	"time"

	probing "github.com/prometheus-community/pro-bing" // Paket für das eigentliche Pingen
)

// PingIP führt einen Ping auf die angegebene IP-Adresse aus und gibt das Ergebnis zurück
func PingIP(ip string) models.PingResult {
	return models.Ping(ip)
}

// PingRange führt einen Ping auf einen Bereich von IP-Adressen aus und gibt alle Ergebnisse zurück
func PingRange(ipRange string) []models.PingResult {
	return models.PingRange(ipRange)
}

// DisplayPingResult gibt das Ergebnis eines einzelnen Pings aus
func DisplayPingResult(result models.PingResult) {
	console.DisplayPingResult(result)
}

// DisplayPingResults gibt die Ergebnisse mehrerer Pings aus
func DisplayPingResults(results []models.PingResult) {
	console.DisplayPingResults(results)
}

// Bang führt das eigentliche Pingen aus
func Bang() {
	// Einlesen von Kommandozeilenargumenten
	timeout := flag.Duration("t", time.Second*100000, "") // Maximale Wartezeit für Antwort
	interval := flag.Duration("i", time.Second, "")       // Abstand zwischen den Ping-Versuchen
	count := flag.Int("c", -1, "")                        // Anzahl der Pings (-1 = unendlich)
	size := flag.Int("s", 24, "")                         // Größe des ICMP-Pakets
	ttl := flag.Int("l", 64, "")                          // Time-to-Live des ICMP-Pakets
	privileged := flag.Bool("privileged", false, "")      // Ausführen des Pings mit root-Rechten
	flag.Usage = func() {                                 // Funktion für die Anzeige der Hilfe
		console.PrintUsage()
	}
	flag.Parse() // Einlesen der Argumente

	if flag.NArg() == 0 { // Keine IP-Adresse angegeben
		console.PrintUsage()
		return
	}

	host := flag.Arg(0)                                                                       // IP-Adresse oder Hostname
	pinger, err := models.NewPinger(host, *count, *size, interval, timeout, ttl, *privileged) // Erstellen des Pinger-Objekts
	if err != nil {
		console.PrintError("Failed to create pinger: %v\n", err)
		return
	}

	// Abfangen des SIGINT-Signals (CTRL-C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop() // Bei Signal das Pingen beenden
		}
	}()

	// Angeben, was bei erfolgreicher Antwort, Antwortduplikaten oder Beendigung des Pingen ausgegeben wird
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

	if err := pinger.Run(); err != nil { // Starten des Pingen
		console.PrintError("Failed to ping target host: %v\n", err)
	}
}
