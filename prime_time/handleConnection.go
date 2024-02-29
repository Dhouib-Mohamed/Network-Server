package prime_time

import (
	"fmt"
	"net"
	"network-challenges/log"
)

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error(fmt.Errorf("error in Closing Connection :%s", err.Error()))
			return
		}
	}(conn)

	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Error(fmt.Errorf("error reading :%s", err.Error()))
			return
		}

		data := buffer[:n]
		handleRequest(data, conn)
	}
}

func handleRequest(data []byte, conn net.Conn) {
	if data[len(data)-1] != 'n' || data[len(data)-2] != '\\' {
		data = append(data, '\\', 'n')
	}
	start := 0
	for i, b := range data {
		if b == '\\' && i+1 < len(data) && data[i+1] == 'n' {
			getResponse(data[start:i], conn)
			start = i + 2
		}
	}
}
