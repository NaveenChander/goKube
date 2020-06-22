package dal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/naveenchander/GoKube/models"
)
	
// IIDMatch ... IIDMatch
type IIDMatch interface {
	SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int)
	CreateIDMatchRequest(IDMatchRequestID uint64, request string) (models.DBErrorTypes, string)
	UpdateIDMatchResponse(IDMatchRequestID uint64, response string) (models.DBErrorTypes, string)
}

// IDMatchSQLDAL ... IDMatchSQLDAL - SQL Call for IDMatch
type IDMatchSQLDAL struct {
	dbServer, dbUser, dbPassword, dbCatalogue string
	dbPort                                    int
}

// IDMatchTestDAL ... IDMatchTestDAL - Mock Call for IDMatch
type IDMatchTestDAL struct {
	dbServer, dbUser, dbPassword, dbCatalogue string
	dbPort                                    int
}

// SetDBVal ... SetDBVal
func (idDal *IDMatchSQLDAL) SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int) {
	idDal.dbServer = dbServer
	idDal.dbUser = dbUser
	idDal.dbPassword = dbPassword
	idDal.dbCatalogue = dbCatalogue
	idDal.dbPort = dbPort
}

// SetDBVal ... SetDBVal
func (idDal *IDMatchTestDAL) SetDBVal(dbServer, dbUser, dbPassword, dbCatalogue string, dbPort int) {
	idDal.dbServer = dbServer
	idDal.dbUser = dbUser
	idDal.dbPassword = dbPassword
	idDal.dbCatalogue = dbCatalogue
	idDal.dbPort = dbPort
}


// CreateIDMatchRequest ... CreateIDMatchRequest
func (idDal *IDMatchSQLDAL) CreateIDMatchRequest(IDMatchRequestID uint64, request string) (models.DBErrorTypes, string) {
	isRollBack := false

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		idDal.dbServer, idDal.dbUser, idDal.dbPassword, idDal.dbPort, idDal.dbCatalogue)

	var err error

	db, err = sql.Open("mssql", connectionString)
	if err != nil {
		return models.DBConectionUnavailable, "CreateIDMatchRequest: Connection unavailable"
	}
	defer db.Close()

	tx, err2 := db.Begin()
	if err2 != nil {
		return models.DBError, "CreateIDMatchRequest: Unable to create Transaction"
	}

	defer func() {
		if isRollBack == false {
			tx.Rollback()
		}
	}()

	query := fmt.Sprintf("EXEC [IDMatch].[usp_IDMatchRequest_add] @IDMatchRequestID = %d, @Request = '%s'", IDMatchRequestID, request)
	log.Println(query)
	_, err3 := tx.Exec(query)
	if err3 != nil {
		tx.Rollback()
		isRollBack = true
		log.Println("CreateIDMatchRequest" + err3.Error())
		return models.DBdmlFailed, err3.Error()
	}
	tx.Commit()

	return models.DBOK, "OK"

}

// UpdateIDMatchResponse ... UpdateIDMatchResponse
func (idDal *IDMatchSQLDAL) UpdateIDMatchResponse(IDMatchRequestID uint64, response string) (models.DBErrorTypes, string) {
	isRollBack := false

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		idDal.dbServer, idDal.dbUser, idDal.dbPassword, idDal.dbPort, idDal.dbCatalogue)

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

	_, err3 := tx.Exec("EXEC [IDMatch].[usp_IDMatchRequest_upd] @IDMatchRequestID = ?, @Response = ?, @UnitTestBypass = 0", IDMatchRequestID, response)
	if err3 != nil {
		tx.Rollback()
		isRollBack = true
		log.Println(err3)
		return models.DBdmlFailed, err3.Error()
	}
	tx.Commit()

	return models.DBOK, "OK"

}
