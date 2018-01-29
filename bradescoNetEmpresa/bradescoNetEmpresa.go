package bradescoNetEmpresa

import (
	"fmt"
	"html"

	"github.com/mundipagg/boleto-api/tmpl"
	"github.com/mundipagg/boleto-api/util"

	"github.com/PMoneda/flow"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/metrics"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/validations"
)

type bankBradescoNetEmpresa struct {
	validate *models.Validator
	log      *log.Log
}

func New() bankBradescoNetEmpresa {
	b := bankBradescoNetEmpresa{
		validate: models.NewValidator(),
		log:      log.CreateLog(),
	}
	b.validate.Push(validations.ValidateAmount)
	b.validate.Push(validations.ValidateExpireDate)
	b.validate.Push(validations.ValidateBuyerDocumentNumber)
	b.validate.Push(validations.ValidateRecipientDocumentNumber)
	b.validate.Push(validations.ValidateBuyerDocumentNumber)
	b.validate.Push(validations.ValidateRecipientDocumentNumber)

	b.validate.Push(bradescoNetEmpresaValidateAgency)
	b.validate.Push(bradescoNetEmpresaValidateAccount)
	b.validate.Push(bradescoNetEmpresaValidateWallet)
	b.validate.Push(bradescoNetEmpresaValidateAgreement)
	return b
}

func (b bankBradescoNetEmpresa) Log() *log.Log {
	return b.log
}

func (b bankBradescoNetEmpresa) RegisterBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	timing := metrics.GetTimingMetrics()
	r := flow.NewFlow()
	serviceURL := config.Get().URLBradescoNetEmpresa
	xmlResponse := getResponseBradescoNetEmpresaXml()
	jsonReponse := getResponseBradescoNetEmpresaJson()
	from := getResponseBradescoNetEmpresa()
	to := getAPIResponseBradescoNetEmpresa()

	bod := r.From("message://?source=inline", boleto, getRequestBradescoNetEmpresa(), tmpl.GetFuncMaps())
	bod.To("logseq://?type=request&url="+serviceURL, b.log)

	err := signRequest(bod)
	if err != nil {
		return models.BoletoResponse{}, err
	}

	duration := util.Duration(func() {
		bod.To(serviceURL, map[string]string{"method": "POST", "insecureSkipVerify": "true"})
	})

	timing.Push("bradesco-netempresa-register-boleto-online", duration.Seconds())

	bod.To("logseq://?type=response&url="+serviceURL, b.log)
	bod.To("transform://?format=xml", xmlResponse, jsonReponse)
	bodyTransform := fmt.Sprintf("%v", bod.GetBody())
	bodyJson := html.UnescapeString(bodyTransform)
	bod.To("set://?prop=body", bodyJson)

	ch := bod.Choice()
	ch.When(flow.Header("status").IsEqualTo("200"))
	ch.To("transform://?format=json", from, to, tmpl.GetFuncMaps())
	ch.Otherwise()
	ch.To("logseq://?type=response&url="+serviceURL, b.log).To("apierro://")

	switch t := bod.GetBody().(type) {
	case string:
		response := util.ParseJSON(t, new(models.BoletoResponse)).(*models.BoletoResponse)
		return *response, nil
	case models.BoletoResponse:
		return t, nil
	}
	return models.BoletoResponse{}, models.NewInternalServerError("MP500", "Erro interno")
}

func signRequest(bod *flow.Flow) error {

	if !config.Get().MockMode {
		bodyToSign := fmt.Sprintf("%v", bod.GetBody())
		signedRequest, err := util.SignRequest(bodyToSign)
		if err != nil {
			return err
		}
		bod.To("set://?prop=body", signedRequest)
	}

	return nil
}

func (b bankBradescoNetEmpresa) ProcessBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	errs := b.ValidateBoleto(boleto)
	if len(errs) > 0 {
		return models.BoletoResponse{Errors: errs}, nil
	}
	return b.RegisterBoleto(boleto)
}

func (b bankBradescoNetEmpresa) ValidateBoleto(boleto *models.BoletoRequest) models.Errors {
	return models.Errors(b.validate.Assert(boleto))
}

func (b bankBradescoNetEmpresa) GetBankNumber() models.BankNumber {
	return models.BradescoNetEmpresa
}
