package main

import (
	"fmt"
	"net"
)

const (
	serverAddr = "localhost:9000"
	bufferSize = 1024
)

func startServer() {
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Server listening on", serverAddr)

	var buffer [bufferSize]byte
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer[:])
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		fmt.Printf("Received: %s from %s\n", string(buffer[:n]), clientAddr)

		_, err = conn.WriteToUDP([]byte("ACK"), clientAddr)
		if err != nil {
			fmt.Println("Error sending ACK:", err)
		}
	}
}

func main() {
	startServer()
}
