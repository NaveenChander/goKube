package api

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/golang/protobuf/ptypes"

	"github.com/naveenchander/GoKube/configuration"
	"github.com/naveenchander/GoKube/core"
	"github.com/naveenchander/GoKube/dal"
	"github.com/naveenchander/GoKube/models"
)

// AddClientCredentials ... Add Client Credentials
func AddClientCredentials(w http.ResponseWriter, r *http.Request) {

	startDate, _ := ptypes.TimestampProto(time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC))
	cc := models.CustomerCredentials{
		CustomerID:   1,
		ClientAPIKey: "testAPIKey",
		Secret:       "Abcd@1234",
		StartDate:    startDate,
	}

	dalDep := dal.CustomerSQLDAL{}
	dalDep.SetDBVal(configuration.DBSERVER, configuration.DBUSER, configuration.DBPASSWORD, configuration.DBCATALOGUE, configuration.DBPORT)

	core.AddClientCredentials(cc, &dalDep)
}

// CreateAuthCode ... CreateAuthCode
func CreateAuthCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["key"]
	secret := params["secret"]

	str := fmt.Sprintf("%s:%s:", key, secret)
	log.Println(str)
	stringToEncode, err := hex.DecodeString(key + ":" + secret)

	if err != nil {
		log.Println(err)
	}

	log.Println(stringToEncode)

	returnValue := base64.StdEncoding.EncodeToString(stringToEncode)
	fmt.Fprintf(w, "Basic "+returnValue)
	w.WriteHeader(200)
}
