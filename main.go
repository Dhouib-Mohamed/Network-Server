package main

import (
	"network-challenges/PrimeTime"
	"network-challenges/SmokeTest"
)

func main() {
	go SmokeTest.Index()
	PrimeTime.Index()
}
