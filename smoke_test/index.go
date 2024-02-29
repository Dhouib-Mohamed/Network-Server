package smoke_test

import (
	"network-challenges/env"
	"network-challenges/server"
)

func Index() {
	server.TCPServer(env.GetSmokeTestPort(), handleConnection)
}
