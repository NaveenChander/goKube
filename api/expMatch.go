package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/naveenchander/GoKube/core"
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
func ExpMatch(release string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var patron models.Patron

		if core.ValidateAuthHeader(r.Header.Get("Authorization")) == false {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorized"))
			return
		}

		errr := json.NewDecoder(r.Body).Decode(&patron)
		if errr != nil {
			log.Fatal("Error while Marshalling :", errr)
		}

		data, err := json.Marshal(patron)
		if err != nil {
			log.Fatal("Error while Marshalling :", err)
		}

		log.Print(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
