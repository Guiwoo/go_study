package dto

import "bulk_upsert/entity"

type Address struct {
	ID         string `json:"id"`
	Person     string `json:"person"`
	Company    string `json:"company"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	Line3      string `json:"line3"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

func (address Address) ToEntity() entity.Address {
	return entity.Address{
		ID:         address.ID,
		Person:     address.Person,
		Company:    address.Company,
		Line1:      address.Line1,
		Line2:      address.Line2,
		Line3:      address.Line3,
		City:       address.City,
		State:      address.State,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		Email:      address.Email,
		Phone:      address.Phone,
	}
}
