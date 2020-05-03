package models

// PreciseIDServer ... PreciseIDServer
type PreciseIDServer struct {
	PIDXMLVersion    string
	Subscriber       Subscriber
	PrimaryApplicant PrimaryApplicant
	Verbose          string
	Vendor           Vendor
	Options          Options
}

// Subscriber ... Subscriber
type Subscriber struct {
	Preamble   string // Preamble = "TRS2"
	OpInitials string // OpInitials = "ES"
	SubCode    int    // SubCode = 1974710
}

// PrimaryApplicant ... PrimaryApplicant
type PrimaryApplicant struct {
	Name              Name
	SSN               string
	CurrentAddress    CurrentAddress
	PreviousAddress   PreviousAddress
	DriverLicense     PrimaryApplicantDriverLicense
	Phones            []PrimaryApplicantPhone
	Employment        Employment
	Age               int
	DOB               string
	YOB               int
	MothersMaidenName string
	EmailAddress      string
}

// Name ... Name
type Name struct {
	Surname string
	First   string
	Middle  string
	Gen     string //  one of the following: SR, JR, 2, II, 3, III, 4, IV
}

// CurrentAddress ... CurrentAddress
type CurrentAddress struct {
	Street string
	City   string
	State  string
	Zip    string
}

// PreviousAddress ... PreviousAddress
type PreviousAddress struct {
	Street string
	City   string
	State  string
	Zip    string
}

// PrimaryApplicantDriverLicense ... PrimaryApplicantDriverLicense
type PrimaryApplicantDriverLicense struct {
	State  string
	Number string
}

// PrimaryApplicantPhone ... PrimaryApplicantPhone
type PrimaryApplicantPhone struct {
	Number string
	Type   string
}

// Employment ... Employment
type Employment struct {
	Company string
	Address string
	City    string
	State   string
	Zip     string
}

// Vendor ... Vendor
type Vendor struct {
	VendorNumber  string // VendorNumber = "000"
	VendorVersion string
}

// Options ... Options
type Options struct {
	BrokerNumber    int
	FreezeKeyPIN    string
	ReferenceNumber string
	ProductOption   int
	DetailRequest   string
	InquiryChannel  string
}
