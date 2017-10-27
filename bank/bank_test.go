package bank

import (
	"testing"

	"github.com/mundipagg/boleto-api/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestShouldExecuteBBStrategy(t *testing.T) {
	assert := func(n models.BankNumber) {
		bank, err := Get(n)
		number := bank.GetBankNumber()
		So(err, ShouldBeNil)
		So(number.IsBankNumberValid(), ShouldBeTrue)
		So(number, ShouldEqual, n)
	}
	Convey("deve-se verificar o retorno da estrategia de cada banco", t, func() {
		assert(models.BancoDoBrasil)
		assert(models.Citibank)
		assert(models.Santander)
		assert(models.Bradesco)
		assert(models.Itau)
		assert(models.Caixa)

		_, err := Get(88)
		So(err, ShouldNotBeNil)
	})

}
