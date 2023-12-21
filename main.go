package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("TCP probe from %s on port %d", conn.RemoteAddr(), conn.LocalAddr().(*net.TCPAddr).Port)
}

func startListener(port int) {
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
		go handleConnection(conn)
	}
}

type statusHandler struct {
	port int
}

func (sh *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP request from %s to port %d", r.RemoteAddr, sh.port)
	log.Printf("%s %s %s", r.Method, r.URL, r.Proto)
	log.Printf("Host: %s", r.Host)
	for name, headers := range r.Header {
		for _, h := range headers {
			log.Printf("%v: %v", name, h)
		}
	}
	response := map[string]string{"status": "OK"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	go startListener(3000)
	go startListener(6000)

	go http.ListenAndServe(":8080", &statusHandler{port: 8080})
	http.ListenAndServe(":9000", &statusHandler{port: 9000})
}