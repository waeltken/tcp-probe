package main

import (
	"log"
	"net"
	"strconv"
)

func handleConnection(conn net.Conn, source string, port int) {
	defer conn.Close()
	log.Printf("TCP probe from %s on port %d", source, port)
}

func startListener(port int, source string) {
	address := ":" + strconv.Itoa(port)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("failed to listen on port %d: %v", port, err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn, source, port)
	}
}

func main() {
	go startListener(3000, "kubelet")
	startListener(6000, "ALB")
}