package server

import (
	"fmt"
	"net"
	"network-challenges/src/utils/log"
)

func TCPServer(port int, handleConnection func(net.Conn)) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Error(fmt.Errorf("error in Creating server :%s", err.Error()))
		return
	}
	log.Info(fmt.Sprintf("server is Listening on Port %d", port))
	defer func(ln net.Listener) {
		if err := ln.Close(); err != nil {
			log.Error(fmt.Errorf("error in Closing server :%s", err.Error()))
			return
		}
	}(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(fmt.Errorf("error in Receiving Request :%s", err.Error()))
			return
		}
		go handleConnection(conn)
	}
}
