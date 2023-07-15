package main

import "fmt"

const (
	Online      = 0
	Offline     = 0
	Maintenance = 0
	Retired     = 3
)

func printServerStatus(servers map[string]int) {
	fmt.Println("\nThere are", len(servers), "servers")

	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("Unhandled server status")
		}
	}

	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are Offline")
	fmt.Println(stats[Maintenance], "servers are Maintenance")
	fmt.Println(stats[Retired], "servers are Retired")
}

func main() {
	servers := []string{
		"darkstar",
		"aiur",
		"omicron",
		"w359",
		"baseline",
	}

	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}

	printServerStatus(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline
	printServerStatus(serverStatus)

	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}
	printServerStatus(serverStatus)
}
