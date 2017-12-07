package itau

import (
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/validations"
)

func itauValidateAccount(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
<<<<<<< HEAD
		err := t.Agreement.IsAccountValid(7)
=======
		err := t.Agreement.IsAccountValid(8)
>>>>>>> abf3eea4d23f5e043d670dae74627bef4aba106b
		if err != nil {
			return err
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}
