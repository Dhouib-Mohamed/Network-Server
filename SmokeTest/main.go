package main

import "net"
import "fmt"

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		println(n)

		data := buffer[:n]
		println(data)
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}
func main() {
	port := 8080
	ln, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		fmt.Errorf("Error in Creating Server")
	}
	fmt.Printf("Server is Listening on Port %d\n", port)
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Errorf("Error in Receiving Request")
			return
		}
		go handleConnection(conn)
	}
}
