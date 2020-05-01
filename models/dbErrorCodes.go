package models

// DBErrorTypes ... DBErrorTypes based on REST Error codes for ease of translation
type DBErrorTypes int

// DBOK ... HTTP OK
// DBInputErrorClientError ... Bad Request
// DBAuthenticationFailed ... Unauthorized
// DBError ...  Internal Server Error
// DBdmlFailed ... Not Implemented
// DBDuplicateRecordFound ... Bad Gateway
// DBConectionUnavailable ... Service Unavailable
//DBTimeout ...  Gateway Timeout

const (
	DBOK                    DBErrorTypes = 200
	DBInputErrorClientError DBErrorTypes = 400
	DBAuthenticationFailed  DBErrorTypes = 401
	DBError                 DBErrorTypes = 500
	DBdmlFailed             DBErrorTypes = 501
	DBDuplicateRecordFound  DBErrorTypes = 502
	DBConectionUnavailable  DBErrorTypes = 503
	DBTimeout               DBErrorTypes = 504
)
