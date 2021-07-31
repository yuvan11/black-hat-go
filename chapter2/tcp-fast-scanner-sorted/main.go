package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	go workers(ports, results)

	go func() {
		for i := 1; i <= 1024; i++ {

			ports <- i
		}
	}()

	for i := 1; i <= 1024; i++ {

		port := <-results

		if port != 0 {
			openports = append(openports, port)

		}

	}

	close(ports)
	close(results)

	// sort the opening port
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

func workers(ports, results chan int) {

	for port := range ports {

		address := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			// closed port
			results <- 0
			continue
		}
		conn.Close()
		results <- port

	}

}
