package dal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/naveenchander/GoKube/configuration"
	"github.com/naveenchander/GoKube/models"
)

var db *sql.DB
var ctx context.Context

// AddClientCredentialsCustomer ... Opens DB Connections and return a connection object
func AddClientCredentialsCustomer(ClientAPIKey string, secret string, customerID int, startDate time.Time, endDate time.Time) (models.DBErrorTypes, string) {

	isRollBack := false

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		configuration.DBSERVER, configuration.DBUSER, configuration.DBPASSWORD, configuration.DBPORT, configuration.DBCATALOGUE)

	var err error

	db, err = sql.Open("mssql", connectionString)
	if err != nil {
		return models.DBConectionUnavailable, "Connection unavailable"
	}
	defer db.Close()

	tx, err2 := db.Begin()
	if err2 != nil {
		return models.DBError, "Unable to create Transaction"
	}

	defer func() {
		if isRollBack == false {
			tx.Rollback()
		}
	}()

	res, err3 := tx.Exec("EXEC [Customer].[usp_CustomerCredentials_add] @ClientAPIKey = ?, @Secret = ?, @CustomerID = ?, @startDate = ?, @endDate = null, @UnitTestBypass = 0", ClientAPIKey, secret, customerID, startDate)
	if err3 != nil {
		tx.Rollback()
		isRollBack = true
		log.Println(err3)
		return models.DBdmlFailed, err3
	}
	tx.Commit()

	return models.DBOK, "OK"
}
