package api

import (
	"net/http"

	"github.com/naveenchander/GoKube/core"
	"github.com/naveenchander/GoKube/models"
)

// AddClientCredentials ... Add Client Credentials
func AddClientCredentials(w http.ResponseWriter, r *http.Request) {

	cc := models.CustomerCredentials{
		CustomerID:   1,
		ClientAPIKey: "testClientAPIKey1234",
		Secret:       core.HashString("abc@123"),
		StartDate:    "02/02/2020",
		EndDate:      "02/02/2030",
	}

	core.AddClientCredentials(cc)
}
