package means_to_an_end

import (
	"network-challenges/src/server"
	"network-challenges/src/utils/env"
)

func Index() {
	server.TCPServer(env.GetMeansEndPort(), handleConnection)
}
