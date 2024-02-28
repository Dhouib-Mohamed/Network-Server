package Server

import (
	"fmt"
	"net"
	"network-challenges/Log"
)

func TCPServer(port int, handleConnection func(net.Conn)) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		Log.Error(fmt.Errorf("error in Creating Server :%s", err.Error()))
		return
	}
	Log.Info(fmt.Sprintf("Server is Listening on Port %d\n", port))
	defer func(ln net.Listener) {
		if err := ln.Close(); err != nil {
			Log.Error(fmt.Errorf("error in Closing Server :%s", err.Error()))
			return
		}
	}(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			Log.Error(fmt.Errorf("error in Receiving Request :%s", err.Error()))
			return
		}
		go handleConnection(conn)
	}
}
