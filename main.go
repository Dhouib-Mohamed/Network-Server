package main

import (
	"network-challenges/src/main/means_to_an_end"
	"network-challenges/src/main/prime_time"
	"network-challenges/src/main/smoke_test"
)

func main() {
	go smoke_test.Index()
	go prime_time.Index()
	means_to_an_end.Index()
}
