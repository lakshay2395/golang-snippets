package factorypattern

import "errors"

const (
	cash = 1
	creditcard = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetMethod(m int) (PaymentMethod,error){
	switch m {
	case cash : return &CashPM{},nil
	case creditcard : return &CreditCardPM{},nil
	default: return nil,errors.New("invalid payment method")
	}
	return nil,errors.New("not implemented")
}

type CashPM struct{}

type CreditCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return "paid in cash"
}

func (c *CreditCardPM) Pay(amount float32) string {
	return "paid in credit card"
}
