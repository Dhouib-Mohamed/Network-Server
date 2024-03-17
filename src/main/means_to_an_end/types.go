package means_to_an_end

import (
	"fmt"
	"network-challenges/src/utils/log"
)

type Price struct {
	Price     int32
	Timestamp int32
}

type Response struct {
	Action  string
	Success bool
	Result  int32
}
type IRequest interface {
	execute(data *[]Price) (Response, error)
}

type InsertRequest struct {
	Action    string
	Price     int32
	Timestamp int32
}

func (r InsertRequest) execute(data *[]Price) (Response, error) {
	if r.Action != "I" || r.Price == 0 || r.Timestamp == 0 {
		log.Error(fmt.Errorf("invalid Request : %v", r))
		return Response{Action: "I", Success: false}, fmt.Errorf("invalid Request : %v", r)
	}
	newPrice := Price{Price: r.Price, Timestamp: r.Timestamp}
	*data = append(*data, newPrice)
	return Response{Action: "I", Success: true}, nil
}

type QueryRequest struct {
	Action string
	start  int32
	end    int32
}

func (r QueryRequest) execute(data *[]Price) (Response, error) {
	if r.Action != "Q" || r.start == 0 || r.end == 0 {
		log.Error(fmt.Errorf("invalid Request : %v", r))
		return Response{Action: "Q", Success: false}, fmt.Errorf("invalid Request : %v", r)
	}
	total := int32(0)
	count := int32(0)
	for _, price := range *data {
		if price.Timestamp >= r.start && price.Timestamp <= r.end {
			total += price.Price
			count++
		}
	}
	var result int32
	if count == 0 {
		result = 0
	} else {
		result = total / count
	}
	return Response{Action: "Q", Success: true, Result: result}, nil
}
