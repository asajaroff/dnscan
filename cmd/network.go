package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// This function should receive a complete FQDN
// and attempt to connect to every mapped port
// https://en.wikipedia.org/wiki/Fully_qualified_domain_name
func attemptConnections()  {
	mappedPorts := make(map[string]int) // Declaring a map of common ports
	mappedPorts["http"] = 80
	mappedPorts["https"] = 443
	//	mappedPorts["ssh"] = 22
	//	mappedPorts["smtp"] = 25

	var targetDomain string = "golang.org"

	for s, i := range mappedPorts {
		tempDomain := fmt.Sprintf("%s:%d", targetDomain, i)
		conn, err := net.Dial("tcp", tempDomain )
		if err != nil {
			log.Fatalf("Error connecting to %s", tempDomain)
		}
		log.Printf("Remote server address: %s",conn.RemoteAddr())
		log.Printf("Local address: %s",conn.LocalAddr())

		log.Printf("Attempting to connect to %s", s )
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Error sending data to server: %s", err)
		}
		log.Println(status)
		conn.Close()
	}
}
