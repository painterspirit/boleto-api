package caixa

//Response focado sna integracao com a Caixa
const registerBoletoResponseCaixa = `{
    {{if (eq .returnCode "1")}}
        "Errors":[{
            "Code":"{{trim .returnCode}}",
            "Message":"{{trim .returnMessage}}"
        }]
    {{else}}
        "DigitableLine": "{{fmtDigitableLine (trim .digitableLine)}}",
        "BarCodeNumber": "{{trim .barcodeNumber}}",
        "Links": [{
            "href":"{{trim .url}}",
            "rel": "pdf",
            "method":"GET"
        }]

        {{if (ne .ourNumber "0") }}        
        , "OurNumber": "{{trim .ourNumber}}"
        {{end}}    
    {{end}}
}
`

func getAPIResponseCaixa() string {
	return registerBoletoResponseCaixa
}
