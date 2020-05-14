package core

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
)

// AddClientCredentials ... AddClientCredentials
func AddClientCredentials(clientCredentials models.CustomerCredentials, d dal.CustomerSQLDAL) (models.ApplicationErrorCodes, string) {

	startDate, err := ptypes.Timestamp(clientCredentials.StartDate)
	if err != nil {
		log.Fatal("Unable convert start Date")
	}
	var endDate time.Time
	if clientCredentials.EndDate != nil {
		endDate, err = ptypes.Timestamp(clientCredentials.EndDate)
		if err != nil {
			log.Fatal("Unable convert end Date")
		}
	}

	if errCode, errValue := d.AddClientCredentialsCustomer(clientCredentials.ClientAPIKey, HashString(clientCredentials.Secret), int(clientCredentials.CustomerID), startDate, endDate); errCode == models.DBOK {
		return models.HTTPOK, "Successfully Added"
	} else {
		switch errCode {
		case models.DBInputErrorClientError:
			return models.HTTPBadRequest, errValue
		default:
			return models.HTTPInternalServerError, errValue
		}
	}
}
