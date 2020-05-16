package core

import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/naveenchander/GoKube/configuration"
	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
)

// ProcessExperian20 .. ProcessExperian20
func ProcessExperian20(incomingRequest string, expDal dal.IExperian) (models.ApplicationErrorCodes, string) {
	var patron models.Patron
	json.Unmarshal([]byte(incomingRequest), &patron)

	data, err := json.Marshal(patron)
	if err != nil {
		log.Println("ProcessExperian20: Error while Marshalling :", err)
		return models.HTTPBadRequest, err.Error()
	}

	if ret, errREG := regexp.Match(`^\d`, []byte(patron.TIN)); errREG != nil || ret == false {
		log.Println("Invalid TIN number :" + patron.TIN)
		return models.HTTPBadRequest, err.Error()
	}

	flakeID := configuration.GetNextFlakeID()
	returnValue, errValue := expDal.CreateExperianRequest(flakeID, incomingRequest)
	if returnValue != models.DBOK {
		return models.HTTPInternalServerError, errValue
	}

	// TODO: Call Exp 20 Outbound

	response := string(data[:])

	returnValue, errValue = expDal.UpdateExperianResponse(flakeID, response)
	if returnValue != models.DBOK {
		return models.HTTPInternalServerError, errValue
	}
	return models.HTTPOK, string(data[:])

}
