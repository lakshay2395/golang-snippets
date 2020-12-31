package factorypattern

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetMethod(cash)
	if err != nil {
		t.Error(err.Error())
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid in cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}
