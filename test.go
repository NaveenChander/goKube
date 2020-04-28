package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
)

func main() {
	startDate, _ := ptypes.TimestampProto(time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC))
	strtDate, err := ptypes.Timestamp(startDate)

	fmt.Println(strtDate)
	if err != nil {
		log.Fatal("Unable convert start Date")
	}

}
