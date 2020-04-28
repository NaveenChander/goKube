package api

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/naveenchander/GoKube/core"
	"github.com/naveenchander/GoKube/models"
)

// AddClientCredentials ... Add Client Credentials
func AddClientCredentials(w http.ResponseWriter, r *http.Request) {

	startDate, _ := ptypes.TimestampProto(time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC))
	cc := models.CustomerCredentials{
		CustomerID:   1,
		ClientAPIKey: "testClientAPIKey1234",
		Secret:       "abc@123",
		StartDate:    startDate,
	}

	core.AddClientCredentials(cc)
}
