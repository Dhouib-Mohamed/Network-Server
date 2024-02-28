package SmokeTest

import (
	"network-challenges/Env"
	"network-challenges/Server"
)

func Index() {
	Server.TCPServer(Env.GetSmokeTestPort(), handleConnection)
}
