package boleto

import (
	"bytes"
	"encoding/base64"
	"errors"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/twooffive"

	"image/png"

	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/tmpl"
)

type HTMLBoleto struct {
	View          models.BoletoView
	ConfigBank    ConfigBank
	Barcode64     string
	Format        string
	DigitableLine string
}

//HTML renderiza HTML do boleto
func HTML(boletoView models.BoletoView, format string) (string, error) {

	if boletoView.Barcode == "" {
		return "", errors.New("boleto not found")
	}
	b := tmpl.New()
	html := HTMLBoleto{
		View:       boletoView,
		ConfigBank: GetConfig(boletoView.Boleto),
		Format:     format,
	}
	bcode, _ := twooffive.Encode(boletoView.Barcode, true)
	orgBounds := bcode.Bounds()
	orgWidth := orgBounds.Max.X - orgBounds.Min.X
	img, _ := barcode.Scale(bcode, orgWidth, 50)
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	html.Barcode64 = base64.StdEncoding.EncodeToString(buf.Bytes())
	html.DigitableLine = textToImage(boletoView.DigitableLine)
	templateBoleto, boletoForm := getTemplateOfBank(boletoView)
	s, err := b.From(html).To(templateBoleto).Transform(boletoForm)
	if err != nil {
		return "", err
	}
	return s, nil
}

func getTemplateOfBank(boletoView models.BoletoView) (string, string) {

	switch boletoView.Boleto.BankNumber {
	case models.Caixa:
		return getTemplateCaixa()
	default:
		return getTemplateDefault()
	}
}
