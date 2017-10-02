package bradesco

import (
	"testing"
	"time"

	"github.com/mundipagg/boleto-api/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBarcodeGenerationBradesco(t *testing.T) {
	//example := "23795796800000001990001250012446693212345670"
	boleto := models.BoletoRequest{}
	boleto.BankNumber = models.Bradesco
	boleto.Agreement = models.Agreement{
		Account: "1234567",
		Agency:  "1",
		Wallet:  25,
	}
	expireDate, _ := time.Parse("02-01-2006", "01-08-2019")
	boleto.Title = models.Title{
		AmountInCents:  199,
		OurNumber:      124466932,
		ExpireDateTime: expireDate,
	}
	bc := getBarcode(boleto)
	Convey("deve-se montar o código de barras do Bradesco", t, func() {
		So(bc.toString(), ShouldEqual, "23795796800000001990001250012446693212345670")
	})
}
