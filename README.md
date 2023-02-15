# Net-Flashbang
Net Flashbang is an open-source network scanner that helps to have a fast overview of a local area network. This tool can help to understand and troubleshoot this network. 

# 1. Zielbestimmung
## 1.1. Richtziel
Das Ziel ist es ein Netzwerkscanner zu entwickeln, der dazu dient schnell und auch spezifisch sich ein
Überblick über ein Netzwerk zu verschaffen. Dabei handelt sich um lokale Netze, bei welchen zum Beil-
spiel support geleistet werden soll, jedoch keine Dokumentation oder Kenntnisse des Netzes vorhanden
ist.
## 1.2. Ziele / Erfolgskriterien
| Ziel-Nr. | Beschreibung | Wertung (ME / KE) |
|----------|--------------|-------------------|
| 1 | Das Pingen und Anzeigen des Ergebnisses einer einzelnen IP-Adresse muss immer schnell und innerhalb einer Sekunde erfolgen können. | ME |
| 2 |Das Pingen und Anzeigen des Ergebnisses eines Subnetzes muss immer schnell und innerhalb 10 Sekunde erfolgen können.| ME |
| 3 | Das Pingen und Anzeigen des Ergebnisses aller RFC1918 definierten Subnetze muss immer schnell und innerhalb 5 Minuten erfolgen können. | ME |
| 4 | Zufügen von Funktionen im Ping Befehl muss immer gewährleistet werden | ME |
| 5 | Die Funktion «-m» oder «--mac» funktioniert immer und gibt die MAC-Adresse aller gepingten IP-Adressen zusätzlich an | ME |
| 6 | Die Funktion «-h» oder «--host» funktioniert immer und gibt den Hostnamen aller gepingten IP-Adressen zusätzlich an | KE |
| 7 | Die Funktion «-o» oder «--os» funktioniert immer und gibt die OS-Version aller gepingten IP-Adressen zusätzlich an, sofern diese auslesbar vom Endgerät ist | KE |
| 8 | Die Funktion «-p» oder «--port» funktioniert immer und gibt alle individuelle Portinformationen aller gepingten IP-Adressen zusätzlich an sofern diese auslesbar vom Endgerät sind | KE |
| 9 | Die Auswertung der Ergebnisse wird in eine dynamische Table fehlerfrei ausgegeben | ME |

Legende:
ME:
Muss erfüllt werden
KE:
Kann erfüll werden