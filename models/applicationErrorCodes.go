package models

// ApplicationErrorCodes ... ApplicationErrorCodes based on REST Error codes for ease of translation
type ApplicationErrorCodes int

const (
	HTTPOK                  ApplicationErrorCodes = 200
	HTTPBadRequest          ApplicationErrorCodes = 400
	HTTPUnauthorized        ApplicationErrorCodes = 401
	HTTPForbidden           ApplicationErrorCodes = 403
	HTTPNotFound            ApplicationErrorCodes = 404
	HTTPNotAcceptable       ApplicationErrorCodes = 406
	HTTPRequestTimeout      ApplicationErrorCodes = 408
	HTTPInternalServerError ApplicationErrorCodes = 500
	HTTPNotImplemented      ApplicationErrorCodes = 501
	HTTPServerUnavailable   ApplicationErrorCodes = 503
)
