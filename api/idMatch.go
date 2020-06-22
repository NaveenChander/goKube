package api

import (
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

func IDMatch(w http.ResponseWriter, r *http.Request) {

	body, errParse := ioutil.ReadAll(r.Body)
	if errParse != nil {
		log.Fatal("Error Parsing body as JSON")
	}

	log.Println("Read body" + string(body[:]))
	requestBody := string(body[:])

	idMatchdal := dal.IDMatchSQLDAL{}
	idMatchdal.SetDBVal(configuration.DBSERVER, configuration.DBUSER, configuration.DBPASSWORD, configuration.DBCATALOGUE, configuration.DBPORT)

	outbound := outbound.IDMatchCall{}
	outbound.SetIDMatchDetails(configuration.EXPURL, configuration.EXPUSERNAME, configuration.EXPPASSWORD)

	returnCode, returnValue := core.ProcessIDMatch(requestBody, &idMatchdal, outbound)
	if returnCode != models.HTTPOK {
		w.WriteHeader(int(returnCode))
		fmt.Fprintf(w, "Error returned :"+returnValue)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(returnValue))
}
