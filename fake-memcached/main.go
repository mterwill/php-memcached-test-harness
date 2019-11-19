package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:11211")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.Println("fake-memcached server started")
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if n >= len(buf) {
			panic(errors.New("Not expecting a message that long"))
		}

		if err != nil {
			if err == io.EOF {
				log.Printf("%s: Closing connection", conn.RemoteAddr().String())
				conn.Close()
				return
			}

			panic(err)
		}

		log.Printf("%s: Got message %q", conn.RemoteAddr().String(), buf[0:n])

		if string(buf[0:n]) == "get foo\r\n" {
			conn.Write([]byte("VALUE foo 0 3\r\nbar\r\nEND\r\n"))
		} else {
			// for any other command, hang.
			// memcached won't send another command without a response, so this
			// will hang until the client closes the connection.
		}
	}
}
