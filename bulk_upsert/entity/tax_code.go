package entity

type TaxCode struct {
	ID   string
	Name string
}

func (_ TaxCode) TableName() string {
	return "tax_code"
}
