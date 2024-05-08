package main

import "fmt"

const (
	Online = iota
	Offline
	Maintenance
	Retired
)

var statusNames = []string{"Online", "Offline", "Maintenance", "Retired"}

func displayServerInfo(servers map[string]int) {
	totalServers := len(servers)
	statusCounts := make(map[int]int)

	for _, status := range servers {
		statusCounts[status]++
	}

	fmt.Println("Number of servers:", totalServers)
	for i, count := range statusCounts {
		fmt.Printf("Number of %s servers: %d\n", statusNames[i], count)
	}
	fmt.Println()
}

func main() {
	servers := map[string]int{
		"darkstar": Online,
		"aiur":     Online,
		"omicron":  Online,
		"w359":     Online,
		"baseline": Online,
	}

	// Display server info
	displayServerInfo(servers)

	// Change server status of `darkstar` to `Retired`
	servers["darkstar"] = Retired

	// Change server status of `aiur` to `Offline`
	servers["aiur"] = Offline

	// Display server info
	displayServerInfo(servers)

	// Change server status of all servers to `Maintenance`
	for server := range servers {
		servers[server] = Maintenance
	}

	// Display server info
	displayServerInfo(servers)
}
