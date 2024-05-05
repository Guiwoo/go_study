package entity

type Vendor struct {
	ID            string `json:"id"`
	Active        bool   `json:"active"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	CompanyName   string `json:"company_name"`
	AccountNumber string `json:"account_number"`
	Website       string `json:"website"`
	InternalNotes string `json:"internal_notes"`
	CustomFields  string `json:"custom_fields"`
	Currency      string `json:"currency"`

	TaxCodeID     string `json:"tax_code"`
	PaymentTermID string `json:"payment_term"`
	AddressID     string `json:"address"`
	ContractID    string `json:"contract"`
}

func (_ Vendor) TableName() string {
	return "vendor"
}
