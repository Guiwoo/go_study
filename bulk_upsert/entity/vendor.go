package entity

type Vendor struct {
	Id            string
	Active        bool
	Name          string
	Description   string
	CompanyName   string
	AccountNumber string
	Website       string
	InternalNotes string
	CustomFields  string
	Currency      string

	TaxCodeId     string
	PaymentTermId string
	AddressId     string
	ContactId     string
}

func (_ Vendor) TableName() string {
	return "vendor"
}
