package entity

type PaymentTerm struct {
	Id   string
	Name string
}

func (_ PaymentTerm) TableName() string {
	return "payment_term"
}
