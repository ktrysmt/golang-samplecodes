package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	user     string
	password string
	port     int
)

func init() {
	flag.StringVar(&user, "user", "", "input your username")
	flag.StringVar(&password, "pass", "", "input your user's password")
	flag.IntVar(&port, "port", 22, "apply the server's standby port number")
}

func main() {
	exitcode := run()
	os.Exit(exitcode)
}

func run() int {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return 2
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout: 5 * time.Second,
	}

	hostport := fmt.Sprintf("%s:%d", flag.Arg(0), port)

	conn, err := ssh.Dial("tcp", hostport, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot connect %v:%v", hostport, err)
		return 1
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open new session: %v", err)
		return 1
	}
	defer session.Close()

	go func() {
		time.Sleep(5 * time.Second)
		conn.Close()
	}()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	err = session.Run(strings.Join(flag.Args()[1:], " "))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		if ee, ok := err.(*ssh.ExitError); ok {
			return ee.ExitStatus()
		}
		return 1
	}

	return 0
}
