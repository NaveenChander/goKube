package dal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/naveenchander/GoKube/configuration"
)

// GetAuthHeader ... Opens DB Connections and return a connection object
func GetAuthHeader(apiKey, secret string) bool {

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		configuration.DBSERVER, configuration.DBUSER, configuration.DBPASSWORD, configuration.DBPORT, configuration.DBCATALOGUE)
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()

	result := ""

	query := fmt.Sprintf("SELECT DISTINCT 'true' FROM [Customer].[CustomerCredentials] WHERE [ClientAPIKey] = '%s' AND [Secret] = '%s' AND EndDate IS NULL", apiKey, secret)
	err = db.QueryRow(query).Scan(&result)
	if err != nil {
		return false
	}

	if result == "true" {
		return true
	}

	return false
}

// GetExperianAuthorized ... GetExperianAuthorized Checks APIKEY and returns if the user is Authorized
func IsAuthorized(apiKey string, serviceName string) bool {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		configuration.DBSERVER, configuration.DBUSER, configuration.DBPASSWORD, configuration.DBPORT, configuration.DBCATALOGUE)

	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Println("Open connection failed:", err.Error())
		return false
	}
	defer db.Close()

	query := fmt.Sprintf("EXEC [Customer].[usp_GetServiceOFferingsByIDOrApikey] @apiKey = '%s'", apiKey)
	rows, errQuery := db.Query(query)

	if errQuery != nil {
		log.Println(errQuery.Error())
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var name, desc  string
		var serviceOfferingID, internalReferenceID int

		if err := rows.Scan(&serviceOfferingID, &name, &desc, &internalReferenceID); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Println(err.Error())
			return false
		}
		if name == serviceName {
			return true
		}
	}
	return false

}
