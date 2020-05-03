package models

// Response ... Response
type Response struct {
	Error        string
	Summary      Summary
	GLBDetail    GLBDetail
	PreciseMatch PreciseMatch
}

// Summary ... Summary
type Summary struct {
	InitialDecision              string
	FinalDecision                string
	CrossReferenceIndicatorsGrid CrossReferenceIndicatorsGrid
}

// GLBDetail ... GLBDetail
type GLBDetail struct {
	FraudShield FraudShield
}

// PreciseMatch ... PreciseMatch
type PreciseMatch struct {
	PreciseMatchScore    string
	PreciseMatchDecision string
	Addresses            []BaseAddress
	Phones               []Phone
	ConsumerID           ConsumerID
	DateOfBirth          DateOfBirth
	DriverLicense        DriverLicense
	ChangeOfAddresses    ChangeOfAddresses
	OFAC                 OFAC
}

// CrossReferenceIndicatorsGrid ... CrossReferenceIndicatorsGrid
type CrossReferenceIndicatorsGrid struct {
	FullNameVerifiesToAddress    string
	FullNameVerifiesToSSN        string
	FullNameVerifiesToDL         string
	FullNameVerifiesToPhone      string
	SurnameOnlyVerifiesToAddress string
	SurnameOnlyVerifiesToSSN     string
	SurnameOnlyVerifiesToDL      string
	SurnameOnlyVerifiesToPhone   string
	AddressVerifiesToFullName    string
	AddressVerifiesToSurnameOnly string
	AddressVerifiesToSSN         string
	AddressVerifiesToDL          string
	AddressVerifiesToPhone       string
	SSNVerifiesToFullName        string
	SSNVerifiesToSurnameOnly     string
	SSNVerifiesToAddress         string
	DLVerifiesToFullName         string
	DLVerifiesToSurnameOnly      string
	DLVerifiesToAddress          string
	PhoneVerifiesToFullName      string
	PhoneVerifiesToSurnameOnly   string
	PhoneVerifiesToAddress       string
}

// FraudShield ... FraudShield
type FraudShield struct {
	Indicator map[string]string
}

// BaseAddress ... BaseAddress
type BaseAddress struct {
	AddressSummary AddressSummary
	AddressDetail  AddressDetail
}

// AddressSummary ... AddressSummary
type AddressSummary struct {
	VerificationResult map[string]string
	Type               map[string]string
	UnitMismatchResult map[string]string
	HighRiskResult     map[string]string
	Counts             AddressCounts
}

// AddressCounts ... AddressCounts
type AddressCounts struct {
	StandardizedAddressReturnCount int
	ResidentialAddressMatchCount   int
	ResidentialAddressReturnCount  int
	HighRiskAddressReturnCount     int
	BusinessAddressMatchCount      int
	BusinessAddressReturnCount     int
}

// AddressDetail ... AddressDetail
type AddressDetail struct {
	StandardizedAddressRcd     StandardizedAddressRcd
	ResidentialAddressRcds     []Rcd
	HighRiskAddressRcds        []BusinessRcd
	BusinessAddressRcds        []BusinessRcd
	HighRiskAddressDescription HighRiskAddressDescription
}

// StandardizedAddressRcd ... StandardizedAddressRcd
type StandardizedAddressRcd struct {
	Surname   string
	FirstName string
	Middle    string
	Address   string
	City      string
	State     string
	ZipCode   string
	ZipPlus4  string
}

// Rcd ... Rcd
type Rcd struct {
	Surname               string
	FirstName             string
	Middle                string
	AliasName             string
	Address               string
	City                  string
	State                 string
	ZipCode               string
	ZipPlus4              string
	AreaCode              int
	Phone                 int
	OtherHouseholdMembers []OtherHouseholdMember
	MonthsAtResidence     int
	ReportedDate          int
	LastUpdatedDate       int
}

// OtherHouseholdMember ... OtherHouseholdMember
type OtherHouseholdMember struct {
	Name string
}

// BusinessRcd ... BusinessRcd
type BusinessRcd struct {
	BusinessName    string
	Surname         string
	FirstName       string
	Middle          string
	Address         string
	City            string
	State           string
	ZipCode         string
	ZipPlus4        string
	AreaCode        int
	Phone           int
	ReportedDate    int
	LastUpdatedDate int
}

// HighRiskAddressDescription ... HighRiskAddressDescription
type HighRiskAddressDescription struct {
	HighRiskDescription string
}

// Phone ... Phone
type Phone struct {
	Summary PhoneSummary
	Detail  PhoneDetail
}

// PhoneSummary ... PhoneSummary
type PhoneSummary struct {
	VerificationResult map[string]string
	Classification     map[string]string
	HighRiskResult     map[string]string
	Counts             PhoneCounts
}

// PhoneCounts ... PhoneCounts
type PhoneCounts struct {
	ResidentialPhoneMatchCount  int
	ResidentialPhoneReturnCount int
	HighRiskPhoneReturnCount    int
	BusinessPhoneMatchCount     int
	BusinessPhoneReturnCount    int
}

// PhoneDetail ... PhoneDetail
type PhoneDetail struct {
	ResidentialPhoneRcds     []Rcd
	HighRiskPhoneRcds        []BusinessRcd
	BusinessPhoneRcds        []BusinessRcd
	HighRiskPhoneDescription HighRiskAddressDescription
}

// ConsumerID ... ConsumerID
type ConsumerID struct {
	Summary ConsumerIDSummary
	Detail  ConsumerIDDetail
}

// ConsumerIDSummary ... ConsumerIDSummary
type ConsumerIDSummary struct {
	VerificationResult map[string]string
	DeceasedResult     map[string]string
	FormatResult       map[string]string
	IssueResult        map[string]string
	IssueState         string
	IssueStartRange    int
	IssueEndRange      int
	Counts             ConsumerIDCounts
}

// ConsumerIDCounts ... ConsumerIDCounts
type ConsumerIDCounts struct {
	ConsumerIDReturnCount int
}

// ConsumerIDDetail ... ConsumerIDDetail
type ConsumerIDDetail struct {
	ConsumerIDRcds []Rcd
}

// DateOfBirth ... DateOfBirth
type DateOfBirth struct {
	Summary DateOfBirthSummary
}

// DateOfBirthSummary ... DateOfBirthSummary
type DateOfBirthSummary struct {
	MatchResult  map[string]string
	MonthOfBirth int
	DayOfBirth   int
	YearOfBirth  int
}

// DriverLicense ... DriverLicense
type DriverLicense struct {
	Summary DriverLicenseSummary
}

// DriverLicenseSummary ... DriverLicenseSummary
type DriverLicenseSummary struct {
	VerificationResult map[string]string
	FormatValidation   map[string]string
}

// ChangeOfAddresses ... ChangeOfAddresses
type ChangeOfAddresses struct {
	Summary ChangeOfAddressSummary
	Detail  ChangeOfAddressDetail
}

// ChangeOfAddressSummary ... ChangeOfAddressSummary
type ChangeOfAddressSummary struct {
	VerificationResult map[string]string
	Counts             ChangeOfAddressCounts
	Detail             ChangeOfAddressDetail
}

// ChangeOfAddressDetail ... ChangeOfAddressDetail
type ChangeOfAddressDetail struct {
	ChangeOfAddressRcds []ChangeOfAddressRcd
}

// ChangeOfAddressRcd ... ChangeOfAddressRcd
type ChangeOfAddressRcd struct {
	Name            string
	AliasName       string
	Address         string
	City            string
	State           string
	ZipCode         string
	ZipPlus4        string
	ReportedDate    string
	LastUpdatedDate string
}

// ChangeOfAddressCounts ... ChangeOfAddressCounts
type ChangeOfAddressCounts struct {
	ChangeOfAddressReturnCount int
}

// OFAC ... OFAC
type OFAC struct {
	Summary OFACSummary
	Detail  OFACDetail
}

// OFACSummary ... OFACSummary
type OFACSummary struct {
	VerificationResult map[string]string
	Counts             OFACCounts
	Detail             OFACDetail
}

// OFACCounts ... OFACCounts
type OFACCounts struct {
	OFACReturnCount int
}

// OFACDetail ... OFACDetail
type OFACDetail struct {
	OFACRecords []string
}
