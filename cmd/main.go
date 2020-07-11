package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main()  {
	mappedPorts := make(map[string]int)
	mappedPorts["http"] = 80
//	mappedPorts["https"] = 443
//	mappedPorts["ssh"] = 22
//	mappedPorts["smtp"] = 25

	var targetDomain string = "golang.org"

	for s, i := range mappedPorts {
		tempDomain := fmt.Sprintf("%s:%d", targetDomain, i)
		conn, err := net.Dial("tcp", tempDomain )
		if err != nil {
			log.Fatalf("Error connecting to %s", tempDomain)
		}
		if s == "http" {
			fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
			status, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Printf(status)
		}
		conn.Close()
	}
}
