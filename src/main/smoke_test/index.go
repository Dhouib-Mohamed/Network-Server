package smoke_test

import (
	"network-challenges/src/server"
	"network-challenges/src/utils/env"
)

func Index() {
	server.TCPServer(env.GetSmokeTestPort(), handleConnection)
}
