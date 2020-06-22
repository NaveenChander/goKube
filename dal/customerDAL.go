package dal

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang/protobuf/ptypes"
	"github.com/naveenchander/GoKube/models"
)

var db *sql.DB
var ctx context.Context

// ICustomer ... ICustomer
type ICustomer interface {
	AddClientCredentialsCustomer(ClientAPIKey string, secret string, customerID int, startDate time.Time, endDate time.Time) (models.DBErrorTypes, string)
	GetClientCredentials(customerID int) (models.DBErrorTypes, []models.CustomerCredentials)
	SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int)
}

// CustomerSQLDAL ... CustomerSQLDAL
type CustomerSQLDAL struct {
	dbServer, dbUser, dbPassword, dbCatalogue string
	dbPort                                    int
}

// CustomerTestDAL ... CustomerTestDAL
type CustomerTestDAL struct {
	dbServer, dbUser, dbPassword, dbCatalogue string
	dbPort                                    int
}

// SetDBVal ... SetDBVal
func (dal CustomerSQLDAL) SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int) {
	dal.dbServer = dbServer
	dal.dbUser = dbUser
	dal.dbPassword = dbPassword
	dal.dbCatalogue = dbCatalogue
	dal.dbPort = dbPort
}

// SetDBVal ... SetDBVal
func (dal CustomerTestDAL) SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int) {
	dal.dbServer = dbServer
	dal.dbUser = dbUser
	dal.dbPassword = dbPassword
	dal.dbCatalogue = dbCatalogue
	dal.dbPort = dbPort
}

// AddClientCredentialsCustomer ... Opens DB Connections and return a connection object
func (dal CustomerSQLDAL) AddClientCredentialsCustomer(ClientAPIKey string, secret string, customerID int, startDate time.Time, endDate time.Time) (models.DBErrorTypes, string) {

	isRollBack := false

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		dal.dbServer, dal.dbUser, dal.dbPassword, dal.dbPort, dal.dbCatalogue)

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

	_, err3 := tx.Exec("EXEC [Customer].[usp_CustomerCredentials_add] @ClientAPIKey = ?, @Secret = ?, @CustomerID = ?, @startDate = ?, @endDate = null, @UnitTestBypass = 0", ClientAPIKey, secret, customerID, startDate)
	if err3 != nil {
		tx.Rollback()
		isRollBack = true
		log.Println(err3)
		return models.DBdmlFailed, err3.Error()
	}
	tx.Commit()

	return models.DBOK, "OK"
}

// GetClientCredentials ... GetClientCredentials by Customer ID
func (dal CustomerSQLDAL) GetClientCredentials(customerID int) (models.DBErrorTypes, []models.CustomerCredentials) {

	result := make([]models.CustomerCredentials, 1000)

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		dal.dbServer, dal.dbUser, dal.dbPassword, dal.dbPort, dal.dbCatalogue)

	var err error

	db, err = sql.Open("mssql", connectionString)
	if err != nil {
		return models.DBConectionUnavailable, nil
	}
	defer db.Close()

	rows, err := db.Query("EXEC [Customer].[usp_CustomerCredentials_sel] @CustomerID = ?", customerID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var _ClientAPIKey, _Secret, _StartDate, _EndDate string
		var _CustomerID int

		err = rows.Scan(&_ClientAPIKey, &_Secret, &_CustomerID, &_StartDate, &_EndDate)
		if err != nil {
			return models.DBError, nil
		}
		layout := "2006-01-02T15:04:05.000Z"
		sd, _ := time.Parse(layout, _StartDate)
		ed, _ := time.Parse(layout, _EndDate)

		stDate, _ := ptypes.TimestampProto(sd)
		enDate, _ := ptypes.TimestampProto(ed)

		v := models.CustomerCredentials{ClientAPIKey: _ClientAPIKey, Secret: _Secret, CustomerID: int32(customerID), StartDate: stDate, EndDate: enDate}

		result = append(result, v)
	}

	return models.DBOK, result
}
