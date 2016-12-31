package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Multi cast server start at 224.0.0.1")

	conn, err := net.Dial("udp", "224.0.0.1:9999")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	start := time.Now()
	wait := 10*time.Second - time.Nanosecond*time.Duration(start.UnixNano()%(10*1000*1000*1000))
	time.Sleep(wait)

	ticker := time.Tick(10 * time.Second)

	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick: ", now.String())
	}
}
