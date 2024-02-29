package prime_time

import (
	"network-challenges/env"
	"network-challenges/server"
)

func Index() {
	server.TCPServer(env.GetPrimeTimePort(), handleConnection)
}
