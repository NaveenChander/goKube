package core

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
)

// AddClientCredentials ... AddClientCredentials
func AddClientCredentials(clientCredentials models.CustomerCredentials, d dal.ICustomer) (models.ApplicationErrorCodes, string) {

	startDate, err := ptypes.Timestamp(clientCredentials.StartDate)
	if err != nil {
		return models.HTTPBadRequest, err.Error()
	}
	var endDate time.Time
	if clientCredentials.EndDate != nil {
		endDate, err = ptypes.Timestamp(clientCredentials.EndDate)
		if err != nil {
			return models.HTTPBadRequest, err.Error()
		}
	}

	errCode, errValue := d.AddClientCredentialsCustomer(clientCredentials.ClientAPIKey, HashString(clientCredentials.Secret), int(clientCredentials.CustomerID), startDate, endDate)

	if errCode == models.DBOK {
		return models.HTTPOK, "Successfully Added"
	}

	switch errCode {
	case models.DBInputErrorClientError:
		return models.HTTPBadRequest, errValue
	default:
		return models.HTTPInternalServerError, errValue
	}

}
