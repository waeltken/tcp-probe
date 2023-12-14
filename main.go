package main

import (
	"log"
	"net"
)

func handleConnection(conn net.Conn, source string, port int) {
	defer conn.Close()
	log.Printf("TCP probe from %s on port %d", source, port)
}

func main() {
	ln1, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Printf("failed to listen on port 3000: %v", err)
	}
	go func() {
		for {
			conn, err := ln1.Accept()
			if err != nil {
				log.Printf("failed to accept connection: %v", err)
			}
			go handleConnection(conn, "kubelet", 3000)
		}
	}()

	ln2, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Printf("failed to listen on port 6000: %v", err)
	}

	for {
		conn, err := ln2.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
		}
		go handleConnection(conn, "ALB", 6000)
	}
}