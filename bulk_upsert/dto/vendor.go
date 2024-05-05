package dto

import "bulk_upsert/entity"

type Vendor struct {
	ID            string      `json:"id"`
	Active        bool        `json:"active"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	CompanyName   string      `json:"company_name"`
	AccountNumber string      `json:"account_number"`
	Website       string      `json:"website"`
	InternalNotes string      `json:"internal_notes"`
	CustomFields  string      `json:"custom_fields"`
	Currency      string      `json:"currency"`
	TaxCode       TaxCode     `json:"tax_code"`
	PaymentTerm   PaymentTerm `json:"payment_term"`
	Address       Address     `json:"address"`
	Contract      Contract    `json:"contract"`
	AllContacts   []Contract  `json:"all_contacts"`
	AllAddresses  []Address   `json:"all_addresses"`
}

func (vendor Vendor) ToEntity() entity.Vendor {
	return entity.Vendor{
		ID:            vendor.ID,
		Active:        vendor.Active,
		Name:          vendor.Name,
		Description:   vendor.Description,
		CompanyName:   vendor.CompanyName,
		AccountNumber: vendor.AccountNumber,
		Website:       vendor.Website,
		InternalNotes: vendor.InternalNotes,
		CustomFields:  vendor.CustomFields,
		Currency:      vendor.Currency,
		TaxCodeID:     vendor.TaxCode.ID,
		PaymentTermID: vendor.PaymentTerm.ID,
		AddressID:     vendor.Address.ID,
		ContractID:    vendor.Contract.ID,
	}
}
