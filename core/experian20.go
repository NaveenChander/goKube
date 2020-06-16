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

	cacheValue, errCache := getCache("Experian" + patron.TIN)
	if errCache != nil {
		log.Println("Data not found in Cache.  Proceeding to request Experian Cache.")

		// TODO: Call Exp 20 Outbound
		cacheValue = string(data[:])
		_ = setCache("Experian"+patron.TIN, cacheValue, 60)
	}

	response := cacheValue

	returnValue, errValue = expDal.UpdateExperianResponse(flakeID, response)
	if returnValue != models.DBOK {
		return models.HTTPInternalServerError, errValue
	}
	return models.HTTPOK, string(data[:])

}

func buildExperianRequest(patron models.Patron) string {

	// var request models.Experian
	var preciseID models.PreciseIDServer
	var phone models.PrimaryApplicantPhone

	preciseID.Verbose = "Y"
	// Subscriber
	preciseID.Subscriber.OpInitials = "ES"
	preciseID.Subscriber.Preamble = "TRS2"
	preciseID.Subscriber.SubCode = 1974710

	// Vendor
	preciseID.Vendor.VendorNumber = "000"

	// Options
	preciseID.Options.ProductOption = 20

	// primary Applicant
	preciseID.PrimaryApplicant.DOB = patron.Dob
	preciseID.PrimaryApplicant.Name.First = patron.Name.First
	preciseID.PrimaryApplicant.Name.Middle = patron.Name.Middle
	preciseID.PrimaryApplicant.Name.Surname = patron.Name.Last
	preciseID.PrimaryApplicant.Name.Gen = patron.Name.Gen
	preciseID.PrimaryApplicant.SSN = patron.TIN
	preciseID.PrimaryApplicant.DOB = patron.Dob

	// Primary Applicant Current Address
	preciseID.PrimaryApplicant.CurrentAddress.Street = patron.Address.Street
	preciseID.PrimaryApplicant.CurrentAddress.State = patron.Address.State
	preciseID.PrimaryApplicant.CurrentAddress.City = patron.Address.City

	// Primary Applicant Drivers License
	preciseID.PrimaryApplicant.DriverLicense.Number = patron.DriversLicense.Number
	preciseID.PrimaryApplicant.DriverLicense.State = patron.DriversLicense.State

	// primary Applicant Phones
	re := regexp.MustCompile(`[^\d]`)
	phone.Number = re.ReplaceAllString(patron.Phone, "")
	append(preciseID.PrimaryApplicant.Phones, phone)

	return ""
}
