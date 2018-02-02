package bank

import (
	"github.com/mundipagg/boleto-api/itau"
	"github.com/mundipagg/boleto-api/models"
)

func getIntegrationItau(boleto models.BoletoRequest) (Bank, error) {
	return itau.New(), nil
}
