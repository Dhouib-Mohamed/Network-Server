package main

import (
	"network-challenges/prime_time"
	"network-challenges/smoke_test"
)

func main() {
	go smoke_test.Index()
	prime_time.Index()
}
