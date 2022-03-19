package telnet

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type flags struct {
	t *string
}

// Telnet structure
type Telnet struct {
	flags
	host string
	port string
	conn net.Conn
}

// NewTelnet telnet's structure constructor
func NewTelnet() *Telnet {
	return &Telnet{
		flags: flags{
			t: flag.String("timeout", "10s", "server connection timeout"),
		},
		host: "",
		port: "",
	}
}

// Execute runs the utility
func (t *Telnet) Execute() error {
	flag.Parse()

	if len(flag.Args()) < 2 {
		return errors.New("telnet: missing host or port")
	}

	t.host = strings.ToLower(flag.Arg(0))
	t.port = flag.Arg(1)
	timeout, _ := time.ParseDuration(*t.flags.t)

	fmt.Printf("telnet: connecting to %s:%s ...\n", t.host, t.port)

	var err error
	t.conn, err = net.DialTimeout("tcp", net.JoinHostPort(t.host, t.port), timeout)
	if err != nil {
		return errors.New("telnet: connection refused")
	}
	defer t.conn.Close()

	fmt.Println("telnet: connection established")

	errors := make(chan error)

	go func() {
		sender := bufio.NewReader(os.Stdin)
		for {
			data, err := sender.ReadString('\n')
			if err != nil || err == io.EOF {
				errors <- err
			}

			t.conn.Write([]byte(data))
		}
	}()

	go func() {
		receiver := bufio.NewReader(t.conn)
		for {
			data, err := receiver.ReadString('\n')
			if err != nil || err == io.EOF {
				errors <- err
			}

			fmt.Fprint(os.Stdout, string(data))
		}
	}()

	exit := make(chan interface{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGQUIT)

	go func() {
		defer close(exit)

		select {
		case <-interrupt:
			fmt.Println("\ntelnet: connection interrupted")
			return
		case <-errors:
			fmt.Println("\ntelnet: connection closed")
			return
		}
	}()

	<-exit

	return nil
}
