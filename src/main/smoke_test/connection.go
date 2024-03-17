package smoke_test

import (
	"net"
	"network-challenges/src/utils/log"
)
import "fmt"

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		log.Fatal(conn.Close())
	}(conn)

	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Error(fmt.Errorf("error reading :%s", err.Error()))
			return
		}
		data := buffer[:n]

		if _, err = conn.Write(data); err != nil {
			log.Error(fmt.Errorf("error writing :%s", err.Error()))
			return
		}
	}
}
