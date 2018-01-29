package bradescoNetEmpresa

const registerBradescoNetEmpresa = `
## Content-Type:application/json
{
    {{if (eq .Buyer.Document.Type "CNPJ")}}
        "nuCPFCNPJ": "{{splitValues .Recipient.Document.Number 0 8}}",    
    {{else}}
         "nuCPFCNPJ": "{{splitValues .Recipient.Document.Number 0 9}}",	
	{{end}}
    
	{{if (eq .Buyer.Document.Type "CNPJ")}}
        "filialCPFCNPJ": "{{splitValues .Recipient.Document.Number 8 12}}",    
    {{else}}
            "filialCPFCNPJ": "0",	
	{{end}}
	
    {{if (eq .Buyer.Document.Type "CNPJ")}}
        "ctrlCPFCNPJ": "{{splitValues .Recipient.Document.Number 12 14}}",
    {{else}}
        "ctrlCPFCNPJ": "{{splitValues .Recipient.Document.Number 9 11}}",	
    {{end}}	
    "cdTipoAcesso": "2",
    "clubBanco": "2269651",
    "cdTipoContrato": "48",
    "nuSequenciaContrato": "0",
    "idProduto": "{{.Agreement.Wallet}}",
    "nuNegociacao": "{{.Agreement.Agency}}0000000{{.Agreement.Account}}",
    "cdBanco": "237",
    "eNuSequenciaContrato": "0",
    "tpRegistro": "1",
    "cdProduto": "0",
    "nuTitulo": "{{.Title.OurNumber}}",
    "nuCliente": "{{.Title.DocumentNumber}}",
	"dtEmissaoTitulo": "{{brDateDelimiterTime .Title.CreateDate "."}}",
    "dtVencimentoTitulo": "{{brDateDelimiter .Title.ExpireDate "."}}",
    "tpVencimento": "0",
    "vlNominalTitulo": "{{.Title.AmountInCents}}",
    "cdEspecieTitulo": "03",
    "tpProtestoAutomaticoNegativacao": "0",
    "prazoProtestoAutomaticoNegativacao": "0",
    "controleParticipante": "",
    "cdPagamentoParcial": "",
    "qtdePagamentoParcial": "0",
    "percentualJuros": "0",
    "vlJuros": "0",
    "qtdeDiasJuros": "0",
    "percentualMulta": "0",
    "vlMulta": "0",
    "qtdeDiasMulta": "0",
    "percentualDesconto1": "0",
    "vlDesconto1": "0",
    "dataLimiteDesconto1": "",
    "percentualDesconto2": "0",
    "vlDesconto2": "0",
    "dataLimiteDesconto2": "",
    "percentualDesconto3": "0",
    "vlDesconto3": "0",
    "dataLimiteDesconto3": "",
    "prazoBonificacao": "0",
    "percentualBonificacao": "0",
    "vlBonificacao": "0",
    "dtLimiteBonificacao": "",
    "vlAbatimento": "0",
    "vlIOF": "0",
    "nomePagador": "{{.Buyer.Name}}",
    "logradouroPagador": "{{.Buyer.Address.Street}}",
    "nuLogradouroPagador": "{{.Buyer.Address.Number}}",
    "complementoLogradouroPagador": "{{.Buyer.Address.Complement}}",
    "cepPagador": "{{splitValues (sanitizeValues .Buyer.Address.ZipCode) 0 5}}",
    "complementoCepPagador": "{{splitValues (sanitizeValues .Buyer.Address.ZipCode) 5 8}}",
    "bairroPagador": "{{.Buyer.Address.District}}",
    "municipioPagador": "{{.Buyer.Address.City}}",
    "ufPagador": "{{.Buyer.Address.StateCode}}",
    {{if (eq .Buyer.Document.Type "CPF")}}
    	"cdIndCpfcnpjPagador": "1",
    {{else}}
        "cdIndCpfcnpjPagador": "2",
    {{end}}
    "nuCpfcnpjPagador": "{{sanitizeValues .Buyer.Document.Number}}",
    "endEletronicoPagador": "",
    "nomeSacadorAvalista": "",
    "logradouroSacadorAvalista": "",
    "nuLogradouroSacadorAvalista": "",
    "complementoLogradouroSacadorAvalista": "",
    "cepSacadorAvalista": "0",
    "complementoCepSacadorAvalista": "0",
    "bairroSacadorAvalista": "",
    "municipioSacadorAvalista": "",
    "ufSacadorAvalista": "",
    "cdIndCpfcnpjSacadorAvalista": "0",
    "nuCpfcnpjSacadorAvalista": "0",
    "endEletronicoSacadorAvalista": ""
}
`

const reponseBradescoNetEmpresaXml = `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <ns2:registrarTituloResponse xmlns:ns2="http://ws.registrotitulo.ibpj.web.bradesco.com.br/">
			<return>{{contentJson}}</return>
        </ns2:registrarTituloResponse>
    </soapenv:Body>
</soapenv:Envelope>
`

const responseBradescoNetEmpresaJson = `{{.contentJson}}`

const reponseBradescoNetEmpresa = `
{
    "cdErro": "{{returnCode}}",
    "msgErro": "{{returnMessage}}",
    "cdBarras": "{{barcodeNumber}}",
    "linhaDigitavel": "{{digitableLine}}"
}
`

func getRequestBradescoNetEmpresa() string {
	return registerBradescoNetEmpresa
}

func getResponseBradescoNetEmpresaXml() string {
	return reponseBradescoNetEmpresaXml
}

func getResponseBradescoNetEmpresaJson() string {
	return responseBradescoNetEmpresaJson
}

func getResponseBradescoNetEmpresa() string {
	return reponseBradescoNetEmpresa
}
