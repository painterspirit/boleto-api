package bradesco

import (
	"github.com/PMoneda/flow"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/metrics"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/tmpl"
	"github.com/mundipagg/boleto-api/util"
	"github.com/mundipagg/boleto-api/validations"
)

type bankBradesco struct {
	validate *models.Validator
	log      *log.Log
}

//New creates a new Bradesco instance
func New() bankBradesco {
	b := bankBradesco{
		validate: models.NewValidator(),
		log:      log.CreateLog(),
	}
	b.validate.Push(validations.ValidateAmount)
	b.validate.Push(validations.ValidateExpireDate)
	b.validate.Push(validations.ValidateBuyerDocumentNumber)
	b.validate.Push(validations.ValidateRecipientDocumentNumber)
	b.validate.Push(bradescoValidateAgency)
	b.validate.Push(bradescoValidateAccount)
	b.validate.Push(bradescoValidateWallet)

	return b
}

//Log retorna a referencia do log
func (b bankBradesco) Log() *log.Log {
	return b.log
}

func (b bankBradesco) RegisterBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	timing := metrics.GetTimingMetrics()
	r := flow.NewFlow()
	serviceURL := config.Get().URLBradesco
	from := getResponseBradesco()
	to := getAPIResponseBradesco()
	bod := r.From("message://?source=inline", boleto, getRequestBradesco(), tmpl.GetFuncMaps())
	bod.To("logseq://?type=request&url="+serviceURL, b.log)
	duration := util.Duration(func() {
		bod.To(serviceURL, map[string]string{"method": "POST", "insecureSkipVerify": "true"})
	})
	timing.Push("bradesco-shopfacil-register-boleto-online", duration.Seconds())
	bod.To("logseq://?type=response&url="+serviceURL, b.log)
	ch := bod.Choice()
	ch.When(flow.Header("status").IsEqualTo("201"))
	ch.To("transform://?format=json", from, to, tmpl.GetFuncMaps())
	ch.To("unmarshall://?format=json", new(models.BoletoResponse))
	ch.When(flow.Header("status").IsEqualTo("200"))
	ch.To("transform://?format=json", from, to, tmpl.GetFuncMaps())
	ch.To("unmarshall://?format=json", new(models.BoletoResponse))
	ch.Otherwise()
	ch.To("logseq://?type=response&url="+serviceURL, b.log).To("apierro://")
	switch t := bod.GetBody().(type) {
	case *models.BoletoResponse:
		return *t, nil
	case error:
		return models.BoletoResponse{}, t
	}
	return models.BoletoResponse{}, models.NewInternalServerError("MP500", "Internal Error")
}

func (b bankBradesco) ProcessBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	errs := b.ValidateBoleto(boleto)
	if len(errs) > 0 {
		return models.BoletoResponse{Errors: errs}, nil
	}
	return b.RegisterBoleto(boleto)
}

func (b bankBradesco) ValidateBoleto(boleto *models.BoletoRequest) models.Errors {
	return models.Errors(b.validate.Assert(boleto))
}

func (b bankBradesco) GetBankNumber() models.BankNumber {
	return models.Bradesco
}
