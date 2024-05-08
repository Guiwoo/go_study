package dto

import (
	"bulk_upsert/entity"
)

type Contact struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Suffix     string `json:"suffix"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Mobile     string `json:"mobile"`
	AltPhone   string `json:"alt_phone"`
	Fax        string `json:"fax"`
}

func (contract Contact) ToEntity() entity.Contact {
	return entity.Contact{
		Id:         contract.ID,
		Title:      contract.Title,
		FirstName:  contract.FirstName,
		MiddleName: contract.MiddleName,
		LastName:   contract.LastName,
		Suffix:     contract.Suffix,
		Role:       contract.Role,
		Email:      contract.Email,
		Phone:      contract.Phone,
		Mobile:     contract.Mobile,
		AltPhone:   contract.AltPhone,
		Fax:        contract.Fax,
	}
}
