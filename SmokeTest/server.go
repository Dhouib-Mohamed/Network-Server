package SmokeTest

import (
	"fmt"
	"net"
)

func Server() {
	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		err := fmt.Errorf("Error in Creating Server :%s", err.Error())
		fmt.Println(err.Error())
	}
	fmt.Printf("Server is Listening on Port %d\n", port)
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			err := fmt.Errorf("Error in Closing Server :%s", err.Error())
			fmt.Println(err.Error())
		}
	}(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			err := fmt.Errorf("Error in Receiving Request :%s", err.Error())
			fmt.Println(err.Error())
			return
		}
		go handleConnection(conn)
	}
}
