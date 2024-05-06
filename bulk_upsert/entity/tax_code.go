package entity

type TaxCode struct {
	Id   string
	Name string
}

func (_ TaxCode) TableName() string {
	return "tax_code"
}
