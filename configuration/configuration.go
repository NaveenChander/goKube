package configuration

import flake "github.com/davidnarayan/go-flake"

var (
	// Release is a semantic version of current build
	Release = "0.0.2"
	// PORT is a configurable via MakeFile
	PORT = "8000"
	// DBSERVER database server name
	DBSERVER = "localhost"
	// DBPORT SQL Server PORT
	DBPORT = 1433
	// DBUSER DB user name
	DBUSER = "sa"
	// DBPASSWORD password
	DBPASSWORD = "abc@123"
	// DBCATALOGUE DB Name
	DBCATALOGUE = "GoKube"
)

// GoFlake ... GoFlake
var GoFlake *flake.Flake

// InitStartup ... Init Global variables
func InitStartup() {
	GoFlake, _ = flake.New()
}

// GetNextFlakeID ... Returns new Flake ID
func GetNextFlakeID() uint64 {
	return uint64(GoFlake.NextId())
}
