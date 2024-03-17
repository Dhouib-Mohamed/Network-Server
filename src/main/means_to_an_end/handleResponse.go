package means_to_an_end

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"network-challenges/src/utils/log"
)

func getResponse(data []byte, conn net.Conn, prices *[]Price) {
	response, err := extractRequest(data)

	if err != nil {
		log.Error(err)
		_, err := conn.Write([]byte("Invalid Request\n"))
		log.Fatal(err)
		log.Fatal(conn.Close())
		return
	}
	result, err := response.execute(prices)

	if err != nil || !result.Success {
		log.Error(err)
		_, err := conn.Write([]byte("Invalid Request\n"))
		log.Fatal(err)
		log.Fatal(conn.Close())
		return
	}

	if result.Action == "Q" && result.Success {
		data := bytes.NewBuffer([]byte{})
		err := binary.Write(data, binary.BigEndian, result.Result)
		if err != nil {
			return
		}
		dst := make([]byte, hex.EncodedLen(len(data.Bytes())))
		hex.Encode(dst, data.Bytes())
		_, err = conn.Write(dst)
		_, err = conn.Write([]byte("\n"))
		if err != nil {
			log.Error(err)
			log.Fatal(conn.Close())
		}
	}
}

func extractRequest(data []byte) (IRequest, error) {
	action := string(data[0])

	if action == "I" {
		return InsertRequest{Action: action, Timestamp: getNumberFromBytes(data[1:5]), Price: getNumberFromBytes(data[5:9])}, nil
	}
	if action == "Q" {
		return QueryRequest{Action: action, start: getNumberFromBytes(data[1:5]), end: getNumberFromBytes(data[5:9])}, nil
	}
	return nil, fmt.Errorf("invalid Request : %v", data)
}

func getNumberFromBytes(data []byte) int32 {
	buf := bytes.NewBuffer(data)
	var number int32
	err := binary.Read(buf, binary.BigEndian, &number)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
