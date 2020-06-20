package core

import (
	"encoding/json"
	"log"
	"regexp"
	"time"

	"github.com/naveenchander/GoKube/configuration"
	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
	"github.com/naveenchander/GoKube/outbound"
)

// ProcessIDMatch .. ProcessIDMatch
func ProcessIDMatch(incomingRequest string, expDal dal.IIDMatch, expOutbound outbound.IIDMatch) (models.ApplicationErrorCodes, string) {
	var patron models.Patron
	json.Unmarshal([]byte(incomingRequest), &patron)

	log.Println("------ Inside Core -----------")
	data, err := json.Marshal(patron)
	if err != nil {
		log.Println("ProcessIDMatch: Error while Marshalling :", err)
		return models.HTTPBadRequest, err.Error()
	}

	if ret, errREG := regexp.Match(`^\d`, []byte(patron.TIN)); errREG != nil || ret == false {
		log.Println("Invalid TIN number :" + patron.TIN)
		return models.HTTPBadRequest, errREG.Error()
	}

	if _, tErr := time.Parse("2006/01/02", patron.Dob); tErr != nil {
		if _, tErr = time.Parse("01/02/2006", patron.Dob); tErr != nil {
			log.Println("Invalid Date format Must be in [yyyy/mm/dd] or [mm/dd/yyyy] :" + patron.Dob)
			return models.HTTPBadRequest, tErr.Error()
		}

	}

	log.Println("Incoming Request -> " + string(data[:]))

	flakeID := configuration.GetNextFlakeID()
	returnValue, errValue := expDal.CreateIDMatchRequest(flakeID, incomingRequest)
	if returnValue != models.DBOK {
		return models.HTTPInternalServerError, errValue
	}

	cacheValue, errCache := getCacheInBytes("IDMatch:" + patron.TIN)
	if errCache != nil {
		log.Println("Data not found in Cache.  Proceeding to request IDMatch Cache.")

		xmlForIDMatch, errXML := buildIDMatchRequest(patron)
		if errXML != nil {
			log.Println("Invalid Data while creating XML  :" + errXML.Error())
			return models.HTTPBadRequest, errXML.Error()
		}

		ret, err := expOutbound.Dial(xmlForIDMatch)
		if err != nil {
			return models.HTTPInternalServerError, err.Error()
		}
		// TODO: Check Returu Value and remove TODO

		cacheValue = ret
		_ = setCacheInBytes("IDMatch:"+patron.TIN, cacheValue, 60)
	} else {
		log.Println("Found in Cache Responding from Cache")

	}
	log.Println(string(cacheValue[:]))

	response := cacheValue

	returnValue, errValue = expDal.UpdateIDMatchResponse(flakeID, string(response[:]))
	if returnValue != models.DBOK {
		return models.HTTPInternalServerError, errValue
	}
	return models.HTTPOK, string(data[:])
}

func buildIDMatchRequest(patron models.Patron) ([]byte, error) {

	return nil, nil
}
