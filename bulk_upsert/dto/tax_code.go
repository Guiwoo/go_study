package dto

import "bulk_upsert/entity"

type TaxCode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (taxCode TaxCode) ToEntity() entity.TaxCode {
	return entity.TaxCode{
		Id:   taxCode.ID,
		Name: taxCode.Name,
	}
}
