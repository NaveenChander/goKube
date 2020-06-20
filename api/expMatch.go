package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/naveenchander/GoKube/configuration"
	"github.com/naveenchander/GoKube/core"
	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
	"github.com/naveenchander/GoKube/outbound"
)

// Home ... Test Function for ExpMatch will be removed soon
func Home(release string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		patronInfo := &models.Patron{
			Name: &models.PatronName{
				First:  "Shankar",
				Last:   "G",
				Middle: "A",
			},
			Address: &models.PatronAddress{
				Street: "10000 S Eastern",
				City:   "Henderson",
				State:  "NV",
				Zip:    "89052-5542",
			},
			DriversLicense: &models.PatronDriversLicense{
				State:  "",
				Number: "",
			},
			Dob:     "01/20/1959",
			Phone:   "1531535151",
			Product: "CentralCredit",
			TIN:     "666724863",
		}

		data, err := json.Marshal(patronInfo)
		if err != nil {
			log.Fatal("Error while Marshalling :", err)
		}

		// fmt.Fprintf(w, "Welcome Home "+release) ---> Need the change Unit Test

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// ExpMatch ... ExpMatch exposed outside that takes patron details
func ExpMatch(w http.ResponseWriter, r *http.Request) {

	// var patron models.Patron
	// if errJSON := json.NewDecoder(r.Body).Decode(&patron); errJSON != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	body, errParse := ioutil.ReadAll(r.Body)
	if errParse != nil {
		log.Fatal("Error Parsing body as JSON")
	}

	log.Println("Read body" + string(body[:]))
	requestBody := string(body[:])

	dal := dal.ExperianSQLDAL{}
	dal.SetDBVal(configuration.DBSERVER, configuration.DBUSER, configuration.DBPASSWORD, configuration.DBCATALOGUE, configuration.DBPORT)

	outbound := outbound.ExperianCall{}
	outbound.SetExperianDetails(configuration.EXPURL, configuration.EXPUSERNAME, configuration.EXPPASSWORD)

	returnCode, returnValue := core.ProcessExperian20(requestBody, dal, outbound)
	if returnCode != models.HTTPOK {
		w.WriteHeader(int(returnCode))
		fmt.Fprintf(w, "Error returned :"+returnValue)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(returnValue))
}
