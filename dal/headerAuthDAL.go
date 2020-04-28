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

	query := fmt.Sprintf("SELECT 'true' FROM dbo.[CustomerCredentials] WHERE [ClientAPIKey] = '%s' AND [Secret] = '%s'", apiKey, secret)
	err = db.QueryRow(query).Scan(&result)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if result == "true" {
		return true
	}

	return false
}
