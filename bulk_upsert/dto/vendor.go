package dto

import "bulk_upsert/entity"

type Vendor struct {
	ID            string       `json:"id"`
	Active        bool         `json:"active"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	CompanyName   string       `json:"company_name"`
	AccountNumber string       `json:"account_number"`
	Website       string       `json:"website"`
	InternalNotes string       `json:"internal_notes"`
	CustomFields  string       `json:"custom_fields"`
	Currency      string       `json:"currency"`
	TaxCode       *TaxCode     `json:"tax_code"`
	PaymentTerm   *PaymentTerm `json:"payment_term"`
	Address       *Address     `json:"address"`
	Contact       *Contact     `json:"contact"`
	AllContacts   []Contact    `json:"all_contacts"`
	AllAddresses  []Address    `json:"all_addresses"`
}

func (vendor Vendor) ToEntity() entity.Vendor {
	v := entity.Vendor{
		Id:            vendor.ID,
		Active:        vendor.Active,
		Name:          vendor.Name,
		Description:   vendor.Description,
		CompanyName:   vendor.CompanyName,
		AccountNumber: vendor.AccountNumber,
		Website:       vendor.Website,
		InternalNotes: vendor.InternalNotes,
		CustomFields:  vendor.CustomFields,
		Currency:      vendor.Currency,
	}

	if vendor.Address != nil {
		v.AddressId = vendor.Address.ID
	}

	if vendor.TaxCode != nil {
		v.TaxCodeId = vendor.TaxCode.ID
	}

	if vendor.PaymentTerm != nil {
		v.PaymentTermId = vendor.PaymentTerm.ID
	}

	if vendor.Contact != nil {
		v.ContactId = vendor.Contact.ID
	}

	return v
}
