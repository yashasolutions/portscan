package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// scanPort attempts to connect to the specified port on the given host and reports the status.
func scanPort(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		fmt.Printf("Port %d is closed on %s.\n", port, host)
		return
	}
	conn.Close()
	fmt.Printf("Port %d is open on %s.\n", port, host)
}

func main() {
	// Set up command-line arguments
	host := flag.String("host", "", "Host to scan")
	ports := flag.String("ports", "", "Comma-separated list of ports to scan")
	flag.Parse()

	// Ensure the host and ports arguments are provided
	if *host == "" || *ports == "" {
		fmt.Println("Usage: go run port_scanner.go -host=<HOST> -ports=<PORTS>")
		return
	}

	// Split the ports string into a slice of integers
	portList := strings.Split(*ports, ",")
	for _, port := range portList {
		portNum, err := strconv.Atoi(strings.TrimSpace(port))
		if err != nil {
			fmt.Printf("Invalid port: %s\n", port)
			continue
		}
		scanPort(*host, portNum)
	}
}

