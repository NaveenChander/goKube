package core

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
	"github.com/naveenchander/GoKube/outbound"
)

// ProcessExperian20 .. ProcessExperian20
func ProcessExperian20(incomingRequest string, expDal dal.IExperian, expOutbound outbound.IExperian) (models.ApplicationErrorCodes, string) {
	var patron models.Patron
	json.Unmarshal([]byte(incomingRequest), &patron)

	log.Println("------ Inside Core -----------")
	data, err := json.Marshal(patron)
	if err != nil {
		log.Println("ProcessExperian20: Error while Marshalling :", err)
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

	flakeID := GetNextFlakeID()
	returnValue, errValue := expDal.CreateExperianRequest(flakeID, incomingRequest)
	if returnValue != models.DBOK {
		return models.HTTPInternalServerError, errValue
	}

	cacheValue, errCache := getCacheInBytes("Experian:" + patron.TIN)
	if errCache != nil {
		log.Println("Data not found in Cache.  Proceeding to request Experian Cache.")

		xmlForExperian, errXML := buildExperianRequest(patron)
		if errXML != nil {
			log.Println("Invalid Data while creating XML  :" + errXML.Error())
			return models.HTTPBadRequest, errXML.Error()
		}

		ret, err := expOutbound.Dial(xmlForExperian)
		if err != nil {
			return models.HTTPInternalServerError, err.Error()
		}
		// TODO: Check Returu Value and remove TODO

		cacheValue = ret
		_ = setCacheInBytes("Experian:"+patron.TIN, cacheValue, 60)
	} else {
		log.Println("Found in Cache Responding from Cache")

	}
	log.Println(string(cacheValue[:]))

	response := cacheValue

	returnValue, errValue = expDal.UpdateExperianResponse(flakeID, string(response[:]))
	if returnValue != models.DBOK {
		return models.HTTPInternalServerError, errValue
	}
	return models.HTTPOK, string(data[:])
}

func buildExperianRequest(patron models.Patron) ([]byte, error) {

	// var request models.Experian
	var experian models.Experian

	var products models.ExperianProducts
	var phone models.PrimaryApplicantPhone

	products.PreciseIDServer.Verbose = "Y"
	products.PreciseIDServer.PIDXMLVersion = "06.00"
	// Subscriber
	products.PreciseIDServer.Subscriber.OpInitials = "ES"
	products.PreciseIDServer.Subscriber.Preamble = "TRS2"
	products.PreciseIDServer.Subscriber.SubCode = "1974710"

	// Vendor
	products.PreciseIDServer.Vendor.VendorNumber = "000"

	// Options
	products.PreciseIDServer.Options.ProductOption = "20"

	// primary Applicant
	products.PreciseIDServer.PrimaryApplicant.DOB = patron.Dob
	products.PreciseIDServer.PrimaryApplicant.Name.First = patron.Name.First
	products.PreciseIDServer.PrimaryApplicant.Name.Middle = patron.Name.Middle
	products.PreciseIDServer.PrimaryApplicant.Name.Surname = patron.Name.Last
	products.PreciseIDServer.PrimaryApplicant.Name.Gen = patron.Name.Gen
	products.PreciseIDServer.PrimaryApplicant.SSN = patron.TIN

	t, tErr := time.Parse("2006/01/02", patron.Dob)
	if tErr == nil {
		products.PreciseIDServer.PrimaryApplicant.DOB = t.Format("01022006")
	} else {
		t, tErr = time.Parse("01/02/2006", patron.Dob)
		products.PreciseIDServer.PrimaryApplicant.DOB = t.Format("01022006")
	}

	// Primary Applicant Current Address
	products.PreciseIDServer.PrimaryApplicant.CurrentAddress.Street = patron.Address.Street
	products.PreciseIDServer.PrimaryApplicant.CurrentAddress.State = patron.Address.State
	products.PreciseIDServer.PrimaryApplicant.CurrentAddress.City = patron.Address.City
	products.PreciseIDServer.PrimaryApplicant.CurrentAddress.Zip = patron.Address.Zip

	// Primary Applicant Drivers License
	products.PreciseIDServer.PrimaryApplicant.DriverLicense.Number = patron.DriversLicense.Number
	products.PreciseIDServer.PrimaryApplicant.DriverLicense.State = patron.DriversLicense.State

	// primary Applicant Phones
	re := regexp.MustCompile(`[^\d]`)
	phone.Number = re.ReplaceAllString(patron.Phone, "")
	products.PreciseIDServer.PrimaryApplicant.Phones.Phone = append(products.PreciseIDServer.PrimaryApplicant.Phones.Phone, phone)
	experian.FraudSolutions.Request.Products = append(experian.FraudSolutions.Request.Products, products)

	returnString, errXML := xml.MarshalIndent(experian, "", " ")

	if errXML != nil {
		return nil, errors.New("Error Parsing XML ->" + errXML.Error())
	}
	return returnString, nil
}
