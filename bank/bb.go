package bank

import (
	"github.com/mundipagg/boleto-api/bb"
	"github.com/mundipagg/boleto-api/models"
)

func getIntegrationBB(boleto models.BoletoRequest) (Bank, error) {
	return bb.New(), nil
}