package models

// Experian ... Experian
type Experian struct {
	FraudSolutions struct {
		Request struct {
			Products []ExperianProducts `xml:"Products,omitempty"`
		} `xml:"Request,omitempty"`
	} `xml:"FraudSolutions,omitempty"`
}

// ExperianProducts ... ExperianProducts
type ExperianProducts struct {
	PreciseIDServer struct {
		PIDXMLVersion string `xml:"PIDXMLVersion,omitempty"`
		Subscriber    struct {
			Preamble   string `xml:"Preamble,omitempty"`   // Preamble = "TRS2"
			OpInitials string `xml:"OpInitials,omitempty"` // OpInitials = "ES"
			SubCode    string `xml:"SubCode,omitempty"`    // SubCode = 1974710
		} `xml:"Subscriber,omitempty"`
		PrimaryApplicant struct {
			Name struct {
				Surname string `xml:"Surname,omitempty"`
				First   string `xml:"First,omitempty"`
				Middle  string `xml:"Middle,omitempty"`
				Gen     string `xml:"Gen,omitempty"` //  one of the following: SR, JR, 2, II, 3, III, 4, IV
			} `xml:"Name,omitempty"`
			SSN            string `xml:"SSN,omitempty"`
			CurrentAddress struct {
				Street string `xml:"Street,omitempty"`
				City   string `xml:"City,omitempty"`
				State  string `xml:"State,omitempty"`
				Zip    string `xml:"Zip,omitempty"`
			} `xml:"CurrentAddress,omitempty"`
			PreviousAddress struct {
				Street string `xml:"Street,omitempty"`
				City   string `xml:"City,omitempty"`
				State  string `xml:"State,omitempty"`
				Zip    string `xml:"Zip,omitempty"`
			} `xml:"PreviousAddress,omitempty"`
			DriverLicense struct {
				State  string `xml:"State,omitempty"`
				Number string `xml:"Number,omitempty"`
			} `xml:"DriverLicense,omitempty"`
			Phones struct {
				Phone []PrimaryApplicantPhone `xml:"Phone,omitempty"`
			} `xml:"Phones,omitempty"`
			Employment struct {
				Company string `xml:"Company,omitempty"`
				Address string `xml:"Address,omitempty"`
				City    string `xml:"City,omitempty"`
				State   string `xml:"State,omitempty"`
				Zip     string `xml:"Zip,omitempty"`
			} `xml:"Employment,omitempty"`
			Age               string `xml:"Age,omitempty"`
			DOB               string `xml:"DOB,omitempty"`
			YOB               string `xml:"YOB,omitempty"`
			MothersMaidenName string `xml:"MothersMaidenName,omitempty"`
			EmailAddress      string `xml:"EmailAddress,omitempty"`
		} `xml:"PrimaryApplicant,omitempty"`
		Verbose string `xml:"Verbose,omitempty"` // Verbose = 'Y"
		Vendor  struct {
			VendorNumber  string `xml:"VendorNumber,omitempty"` // VendorNumber = "000"
			VendorVersion string `xml:"VendorVersion,omitempty"`
		} `xml:"Vendor,omitempty"`
		Options struct {
			BrokerNumber    string `xml:"BrokerNumber,omitempty"`
			FreezeKeyPIN    string `xml:"FreezeKeyPIN,omitempty"`
			ReferenceNumber string `xml:"ReferenceNumber,omitempty"`
			ProductOption   string `xml:"ProductOption,omitempty"` // ProductOption = 20
			DetailRequest   string `xml:"DetailRequest,omitempty"`
			InquiryChannel  string `xml:"InquiryChannel,omitempty"`
		} `xml:"Options,omitempty"`
	} `xml:"PreciseIDServer,omitempty"`
}

// PrimaryApplicantPhone ... PrimaryApplicantPhone
type PrimaryApplicantPhone struct {
	Number string `xml:"Number,omitempty"`
	Type   string `xml:"Type,omitempty"`
}
