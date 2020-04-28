package models

// DBErrorTypes ... DBErrorTypes based on REST Error codes for ease of translation
type DBErrorTypes int

const (
	// DBOK ... HTTP OK
	DBOK DBErrorTypes = 200
	// DBInputErrorClientError ... Bad Request
	DBInputErrorClientError DBErrorTypes = 400
	// DBAuthenticationFailed ... Unauthorized
	DBAuthenticationFailed DBErrorTypes = 401
	// DBError ...  Internal Server Error
	DBError DBErrorTypes = 500
	//DBdmlFailed ... Not Implemented
	DBdmlFailed DBErrorTypes = 501
	// DBDuplicateRecordFound ... Bad Gateway
	DBDuplicateRecordFound DBErrorTypes = 502
	// DBConectionUnavailable ... Service Unavailable
	DBConectionUnavailable DBErrorTypes = 503
	//DBTimeout ...  Gateway Timeout
	DBTimeout DBErrorTypes = 504
)
