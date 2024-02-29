package prime_time

import (
	"encoding/json"
	"fmt"
	"net"
	"network-challenges/log"
)

type IResponse interface {
	validateResponse() bool
	getResponse() (string, error)
}
type Response2 struct {
	Method string `json:"method"`
	Number *int   `json:"number,omitempty"`
}

func (r Response2) validateResponse() bool {
	return r.Method == "isPrime" && r.Number != nil
}

func (r Response2) getResponse() (string, error) {
	return fmt.Sprintf("Response : { method : %s , number : %d }", r.Method, *r.Number), nil
}

type Response1 struct {
	Method string `json:"method"`
	Prime  *bool  `json:"prime,omitempty"`
	Number *int   `json:"number,omitempty"`
}

func (r Response1) validateResponse() bool {
	if !(r.Method == "isPrime" && r.Prime != nil && r.Number != nil) {
		return false
	}
	return checkPrime(*r.Number) == *r.Prime
}

func checkPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (r Response1) getResponse() (string, error) {
	return fmt.Sprintf("Response : { method : %s , number : %d , prime : %t }", r.Method, *r.Number, *r.Prime), nil
}

func getResponse(data []byte, conn net.Conn) {
	res1 := Response1{}
	err1 := validateResponse(data, &res1)
	if err1 == nil {
		result, err := res1.getResponse()
		if err != nil {
			log.Error(err)
			log.Fatal(conn.Close())
			return
		}
		_, err = conn.Write([]byte(result + "\n"))
		log.Fatal(err)
		return
	}
	if err1.Error() == "invalid Response" {
		log.Error(fmt.Errorf("error 1: %s", err1.Error()))
		_, err := conn.Write([]byte("Invalid Request\n"))
		log.Fatal(err)
		log.Fatal(conn.Close())
		return
	}
	res2 := Response2{}
	err2 := validateResponse(data, &res2)
	if err2 == nil {
		result, err := res2.getResponse()
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal(conn.Close())
			return
		}
		_, err = conn.Write([]byte(result + "\n"))
		log.Fatal(err)
		return
	}
	log.Error(fmt.Errorf("Error 1: %s \nError 2: %s", err1.Error(), err2.Error()))
	_, err := conn.Write([]byte("Invalid Request\n"))
	log.Fatal(err)
	log.Fatal(conn.Close())
}

func validateResponse(data []byte, res IResponse) error {
	if err := json.Unmarshal(data, res); err != nil {
		return fmt.Errorf("error in Unmarshalling Response :%s", err.Error())
	}
	if !res.validateResponse() {
		return fmt.Errorf("invalid Response")
	}
	return nil
}
