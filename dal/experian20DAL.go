package dal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/naveenchander/GoKube/models"
)

// IExperian ... IExperian
type IExperian interface {
	SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int)
	CreateExperianRequest(experianRequestID uint64, request string) (models.DBErrorTypes, string)
	UpdateExperianResponse(experianRequestID uint64, response string) (models.DBErrorTypes, string)
}

// ExperianSQLDAL ... ExperianSQLDAL - SQL Call for Experian
type ExperianSQLDAL struct {
	dbServer, dbUser, dbPassword, dbCatalogue string
	dbPort                                    int
}

// ExperianTestDAL ... ExperianTestDAL - Mock Call for Experian
type ExperianTestDAL struct {
	dbServer, dbUser, dbPassword, dbCatalogue string
	dbPort                                    int
}

// SetDBVal ... SetDBVal
func (expDal *ExperianSQLDAL) SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int) {
	expDal.dbServer = dbServer
	expDal.dbUser = dbUser
	expDal.dbPassword = dbPassword
	expDal.dbCatalogue = dbCatalogue
	expDal.dbPort = dbPort
}

// SetDBVal ... SetDBVal
func (expDal *ExperianTestDAL) SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int) {
	expDal.dbServer = dbServer
	expDal.dbUser = dbUser
	expDal.dbPassword = dbPassword
	expDal.dbCatalogue = dbCatalogue
	expDal.dbPort = dbPort
}

// CreateExperianRequest ... CreateExperianRequest
func (expDal *ExperianSQLDAL) CreateExperianRequest(experianRequestID uint64, request string) (models.DBErrorTypes, string) {
	isRollBack := false

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		expDal.dbServer, expDal.dbUser, expDal.dbPassword, expDal.dbPort, expDal.dbCatalogue)

	log.Println("Connection String --> " + connectionString)
	var err error

	db, err = sql.Open("mssql", connectionString)
	if err != nil {
		return models.DBConectionUnavailable, "CreateExperianRequest: Connection unavailable"
	}
	defer db.Close()

	tx, err2 := db.Begin()
	if err2 != nil {
		return models.DBError, "CreateExperianRequest: Unable to create Transaction"
	}

	defer func() {
		if isRollBack == false {
			tx.Rollback()
		}
	}()

	query := fmt.Sprintf("EXEC [experian].[usp_ExperianRequest_add] @ExperianRequestID = %d, @Request = '%s'", experianRequestID, request)
	log.Println(query)
	_, err3 := tx.Exec(query)
	if err3 != nil {
		tx.Rollback()
		isRollBack = true
		log.Println("CreateExperianRequest: " + err3.Error())
		return models.DBdmlFailed, err3.Error()
	}
	tx.Commit()

	return models.DBOK, "OK"

}

// UpdateExperianResponse ... UpdateExperianResponse
func (expDal *ExperianSQLDAL) UpdateExperianResponse(experianRequestID uint64, response string) (models.DBErrorTypes, string) {
	isRollBack := false

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		expDal.dbServer, expDal.dbUser, expDal.dbPassword, expDal.dbPort, expDal.dbCatalogue)

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

	_, err3 := tx.Exec("EXEC [experian].[usp_ExperianRequest_upd] @experianRequestID = ?, @Response = ?, @UnitTestBypass = 0", experianRequestID, response)
	if err3 != nil {
		tx.Rollback()
		isRollBack = true
		log.Println(err3)
		return models.DBdmlFailed, err3.Error()
	}
	tx.Commit()

	return models.DBOK, "OK"

}
