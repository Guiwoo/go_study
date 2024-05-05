package entity

type PaymentTerm struct {
	ID   string
	Name string
}

func (_ PaymentTerm) TableName() string {
	return "payment_term"
}
