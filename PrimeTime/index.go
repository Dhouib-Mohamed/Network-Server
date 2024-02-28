package PrimeTime

import (
	"network-challenges/Env"
	"network-challenges/Server"
)

func Index() {
	Server.TCPServer(Env.GetPrimeTimePort(), handleConnection)
}
