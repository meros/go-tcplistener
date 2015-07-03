package tcplistener

import (
	"net"
	"strconv"
)

func Accept(port int, connchan chan net.Conn) error {

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))

	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		connchan <- conn
	}

	return nil
}
