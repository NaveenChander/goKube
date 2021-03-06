package models

//  IDMatchPatron ... IDMatchPatron
type IDMatchPatron struct {
	Name struct {
		First  string
		Last   string
		Middle string
		Gen    string
	}
	Address struct {
		Street string
		City   string
		State  string
		Zip    string
	}
	DriversLicense struct {
		State  string
		Number string
	}
	Dob      string
	Phone    string
	Product  string
	PatronId string
	Tin      string
	Sid      string
}
