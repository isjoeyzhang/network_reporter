package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func getEtherInterfaces() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	for _, i := range interfaces {
		addresses, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, a := range addresses {
			if strings.Contains(i.Name, "en") {
				switch v := a.(type) {
				case *net.IPAddr:
					//fmt.Println(i.Name)
					result = append(result, i.Name)
				case *net.IPNet:
					//fmt.Println(i.Name, v, v.IP)
					result = append(result, i.Name)
					result = append(result, v.IP.String())
				}
			}
		}
	}

	return result, nil
}

func main() {
	done := make(chan struct{})
	defer close(done)

	run(done)
}

func run(done <-chan struct{}) {
	i := 1

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	serverUri := os.Getenv("NETWORK_REPORT_URL")
	if len(serverUri) < 1 {
		log.Println("need config server uri NETWORK_SERVER_URL")
		return
	}
	for {
		// Do something here
		r, err := getEtherInterfaces()
		if err != nil {
			return
		}
		fmt.Println(r)
		request(serverUri, fmt.Sprintf("%v", r))
		log.Println(i, "done")

		select {
		case <-done:
			log.Println("program finish")
			return
		case <-ticker.C:
			log.Println("redo")
		}
	}
}
