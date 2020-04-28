package core

import (
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
)

// AddClientCredentials ... AddClientCredentials
func AddClientCredentials(clientCredentials models.CustomerCredentials) (error, string) {

	clientCredentials.Secret = HashString(clientCredentials.Secret)

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

	if errCode, errValue := dal.AddClientCredentialsCustomer(clientCredentials.ClientAPIKey, HashString(clientCredentials.Secret), int(clientCredentials.CustomerID), startDate, endDate), errCode == models.DBOK{
		return nil, "Successfully Added"
	}
	else{
		switch errCode {
		case DBInputErrorClientError:
			return 
			
		}
	}

}
