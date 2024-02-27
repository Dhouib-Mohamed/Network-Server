package SmokeTest

import "net"
import "fmt"

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			err := fmt.Errorf("Error reading :%s", err.Error())
			fmt.Println(err.Error())
			return
		}

		println(n)

		data := buffer[:n]
		println(data)
		_, err = conn.Write(data)
		if err != nil {
			err := fmt.Errorf("Error writing :%s", err.Error())
			fmt.Println(err.Error())
			return
		}
	}
}
