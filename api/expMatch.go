package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/naveenchander/GoKube/core"
	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
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
	var patron models.Patron

	if errJSON := json.NewDecoder(r.Body).Decode(&patron); errJSON != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(patron)
	if err != nil {
		log.Fatal("Error while Marshalling :", err)
	}

	if ret, errREG := regexp.Match(`^\d`, []byte(patron.TIN)); errREG != nil || ret == false {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid TIN number :"+patron.TIN)
		return
	}

	dal := dal.Experian{}

	core.ProcessExperian20(patron, dal)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
