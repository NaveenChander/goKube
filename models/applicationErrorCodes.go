package models

// ApplicationErrorCodes ... ApplicationErrorCodes based on REST Error codes for ease of translation
type ApplicationErrorCodes int

const (
	HTTPOK         ApplicationErrorCodes = 200
	HTTPBadRequest ApplicationErrorCodes = 400
)
