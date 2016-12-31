package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("test")

	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received from %v: %v\n", remoteAddr, string(buffer[:length]))

		_, err = conn.WriteTo([]byte("Hey client by server"), remoteAddr)

		if err != nil {
			panic(err)
		}

	}

}
