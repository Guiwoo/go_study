package dto

import "bulk_upsert/entity"

type PaymentTerm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (paymentTerm PaymentTerm) ToEntity() entity.PaymentTerm {
	return entity.PaymentTerm{
		ID:   paymentTerm.ID,
		Name: paymentTerm.Name,
	}
}
