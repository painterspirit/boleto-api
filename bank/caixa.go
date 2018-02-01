package bank

import (
	"github.com/mundipagg/boleto-api/caixa"
	"github.com/mundipagg/boleto-api/models"
)

func getIntegrationCaixa(boleto models.BoletoRequest) (Bank, error) {
	return caixa.New(), nil
}
