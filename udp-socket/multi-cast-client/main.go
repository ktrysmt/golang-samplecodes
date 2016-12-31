package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("multicast client started...")

	addr, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")

	if err != nil {
		panic(err)
	}

	listener, err := net.ListenMulticastUDP("udp", nil, addr)

	defer listener.Close()

	buf := make([]byte, 100)

	for {
		len, remoteAddr, err := listener.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("server %v\n", remoteAddr)

		fmt.Printf("Now    %v\n", string(buf[:len]))
	}
}
