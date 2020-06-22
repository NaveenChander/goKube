package configuration

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
	// EXPCachePeriod EXP MATCH Cache Period in Seconds
	EXPCachePeriod = "60"
	// EXPURL ... EXPURL
	EXPURL = ""
	// EXPUSERNAME ... EXPUSERNAME
	EXPUSERNAME = ""
	// EXPPASSWORD ... EXPPASSWORD
	EXPPASSWORD = ""
	// CryptoKey ... Crypto Key
	CryptoKey = "1709675319508246528"
)
