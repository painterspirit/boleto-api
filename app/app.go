package app

import (
	"fmt"
	"os"

	"github.com/PMoneda/flow"
	"github.com/mundipagg/boleto-api/api"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/env"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/mock"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/util"
)

//Params this struct contains all execution parameters to run application
type Params struct {
	DevMode    bool
	MockMode   bool
	DisableLog bool
}

//NewParams returns new Empty pointer to ExecutionParameters
func NewParams() *Params {
	return new(Params)
}

//Run starts boleto api Application
func Run(params *Params) {
	env.Config(params.DevMode, params.MockMode, params.DisableLog)

	//robot.GoRobots()

	if config.Get().MockMode {
		go mock.Run("9091")
	}

	installLog()
	api.InstallRestAPI()

}

func installLog() {
	err := log.Install()
	if err != nil {
		fmt.Println("Log SEQ Fails")
		os.Exit(-1)
	}
}

func installflowConnectors() {
	flow.RegisterConnector("logseq", util.SeqLogConector)
	flow.RegisterConnector("apierro", models.BoletoErrorConector)
	flow.RegisterConnector("tls", util.TlsConector)
}

func configFlags(devMode, mockMode, disableLog bool) {
	if devMode {
		os.Setenv("INFLUXDB_HOST", "http://localhost")
		os.Setenv("INFLUXDB_PORT", "8086")
		os.Setenv("API_PORT", "3000")
		os.Setenv("API_VERSION", "0.0.1")
		os.Setenv("PDF_API", "http://localhost:8089/topdf")
		os.Setenv("ENVIROMENT", "Development")
		os.Setenv("SEQ_URL", "http://localhost:5341")
		os.Setenv("SEQ_API_KEY", "4jZzTybZ9bUHtJiPdh6")
		os.Setenv("ENABLE_REQUEST_LOG", "false")
		os.Setenv("ENABLE_PRINT_REQUEST", "true")
		os.Setenv("URL_BB_REGISTER_BOLETO", "https://cobranca.homologa.bb.com.br:7101/registrarBoleto")
		os.Setenv("URL_BB_TOKEN", "https://oauth.hm.bb.com.br/oauth/token")
		os.Setenv("CAIXA_ENV", "SGCBS01D")
		os.Setenv("URL_CAIXA", "https://des.barramento.caixa.gov.br/sibar/ManutencaoCobrancaBancaria/Boleto/Externo")
		os.Setenv("URL_CITI", "https://citigroupsoauat.citigroup.com/comercioeletronico/registerboleto/RegisterBoletoSOAP")
		os.Setenv("URL_CITI_BOLETO", "https://ebillpayer.uat.brazil.citigroup.com/ebillpayer/jspInformaDadosConsulta.jsp")
		os.Setenv("APP_URL", "http://localhost:3000/boleto")
		os.Setenv("ELASTIC_URL", "http://localhost:9200")
		os.Setenv("MONGODB_URL", "localhost:27017")
		os.Setenv("MONGODB_USER", "")
		os.Setenv("MONGODB_PASSWORD", "")
		os.Setenv("BOLETO_JSON_STORE", "/home/philippe/boletodb/upMongo")
		os.Setenv("CERT_BOLETO_CRT", "C:\\cert_boleto_api\\certificate.crt")
		os.Setenv("CERT_BOLETO_KEY", "C:\\cert_boleto_api\\mundi.key")
		os.Setenv("CERT_BOLETO_CA", "C:\\cert_boleto_api\\ca-cert.ca")
		os.Setenv("URL_SANTANDER_TICKET", "https://ymbdlb.santander.com.br/dl-ticket-services/TicketEndpointService")
		os.Setenv("URL_SANTANDER_REGISTER", "https://ymbcash.santander.com.br/ymbsrv/CobrancaEndpointService")
		os.Setenv("URL_BRADESCO_SHOPFACIL", "https://homolog.meiosdepagamentobradesco.com.br/api/transacao")
		os.Setenv("ITAU_ENV", "1")
		os.Setenv("SANTANDER_ENV", "T")
		os.Setenv("URL_ITAU_REGISTER", "https://gerador-boletos.itau.com.br/router-gateway-app/public/codigo_barras/registro")
		os.Setenv("URL_ITAU_TICKET", "https://oauth.itau.com.br/identity/connect/token")
		if mockMode {
			os.Setenv("URL_BB_REGISTER_BOLETO", "http://localhost:9091/registrarBoleto")
			os.Setenv("URL_BB_TOKEN", "http://localhost:9091/oauth/token")
			os.Setenv("URL_CAIXA", "http://localhost:9091/caixa/registrarBoleto")
			os.Setenv("URL_CITI", "http://localhost:9091/citi/registrarBoleto")
			os.Setenv("URL_SANTANDER_TICKET", "tls://localhost:9091/santander/get-ticket")
			os.Setenv("URL_SANTANDER_REGISTER", "tls://localhost:9091/santander/register")
			os.Setenv("URL_BRADESCO_SHOPFACIL", "http://localhost:9091/bradescoshopfacil/registrarBoleto")
			os.Setenv("URL_ITAU_TICKET", "http://localhost:9091/itau/gerarToken")
			os.Setenv("URL_ITAU_REGISTER", "http://localhost:9091/itau/registrarBoleto")
		}
	}
	config.Install(mockMode, devMode, disableLog)
}
