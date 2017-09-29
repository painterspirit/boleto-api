package bradesco

import (
	"errors"

	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/validations"
)

func bradescoValidateAgency(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		err := t.Agreement.IsAgencyValid()
		if err != nil {
			return err
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}

func bradescoValidateAccount(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if t.Agreement.Account == "" {
			return errors.New("a conta deve ser preenchida")
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}

func bradescoValidateWallet(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if t.Agreement.Wallet == 0 {
			return errors.New("a wallet deve ser maior que 0")
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}
