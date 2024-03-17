package prime_time

import (
	"network-challenges/src/server"
	"network-challenges/src/utils/env"
)

func Index() {
	server.TCPServer(env.GetPrimeTimePort(), handleConnection)
}
