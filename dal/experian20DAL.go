package dal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/naveenchander/GoKube/configuration"
)

// ExperianSQLDAL ... ExperianSQLDAL - SQL Call for Experian
type ExperianSQLDAL struct{}

// ExperianTestDAL ... ExperianTestDAL - Mock Call for Experian
type ExperianTestDAL struct{}

// GetExperianAuthorized ... GetExperianAuthorized Checks APIKEY and returns if the user is Authorized
func (dal ExperianSQLDAL) GetExperianAuthorized(apiKey) bool {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		configuration.DBSERVER, configuration.DBUSER, configuration.DBPASSWORD, configuration.DBPORT, configuration.DBCATALOGUE)

	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()

	result := ""

	query := fmt.Sprintf("SELECT 'true' FROM [Customer].[CustomerCredentials] WHERE [ClientAPIKey] = '%s' AND [Secret] = '%s' AND EndDate IS NULL", apiKey, secret)
	err = db.QueryRow(query).Scan(&result)

	if err != nil {
		return false
	}

	if result == "true" {
		return true
	}

	return false

}
