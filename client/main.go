package main

import (
	"fmt"
	"net"
	"time"
)

const (
	serverAddr = "localhost:9000"
	bufferSize = 1024
	timeout    = 2 * time.Second
)

func sendReliableMessage(conn *net.UDPConn, addr *net.UDPAddr, message string) error {
	var ackBuffer [bufferSize]byte
	for {
		_, err := conn.Write([]byte(message))
		if err != nil {
			return fmt.Errorf("failed to send message: %v", err)
		}

		conn.SetReadDeadline(time.Now().Add(timeout))
		n, _, err := conn.ReadFromUDP(ackBuffer[:])
		if err == nil {
			message := string(ackBuffer[:n])
			fmt.Println("Message received:", message)
			break
		}

		fmt.Println("Resending message...")
	}
	return nil
}

func startClient() {
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := "Hello, Reliable UDP!"
	err = sendReliableMessage(conn, addr, message)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	startClient()
}
