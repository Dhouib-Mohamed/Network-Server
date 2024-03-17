package means_to_an_end

import (
	"encoding/hex"
	"fmt"
	"net"
	"network-challenges/src/utils/log"
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

	prices := make([]Price, 0)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Error(fmt.Errorf("error reading :%s", err.Error()))
			return
		}

		data := buffer[:n-2]

		req := make([]byte, hex.DecodedLen(len(data)))
		_, err = hex.Decode(req, data)
		if err != nil {
			log.Error(fmt.Errorf("error in Decoding Request :%s", err.Error()))
			return
		}
		handleRequest(req, conn, &prices)
	}
}

func handleRequest(data []byte, conn net.Conn, prices *[]Price) {
	for i := 0; i < len(data)/9; i++ {
		getResponse(data[i*9:(i+1)*9], conn, prices)
	}
}
