package SmokeTest

import (
	"net"
	"network-challenges/Log"
)
import "fmt"

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		Log.Fatal(conn.Close())
	}(conn)

	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log.Error(fmt.Errorf("error reading :%s", err.Error()))
			return
		}
		data := buffer[:n]

		if _, err = conn.Write(data); err != nil {
			Log.Error(fmt.Errorf("error writing :%s", err.Error()))
			return
		}
	}
}
