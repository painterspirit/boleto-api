package itau

import (
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/validations"
)

func itauValidateAccount(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		err := t.Agreement.IsAccountValid(8)
		if err != nil {
			return err
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}
