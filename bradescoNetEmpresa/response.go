package bradescoNetEmpresa

var apiResponse = `
{
	{{if eq .returnCode "0"}}
	   "DigitableLine": "{{.digitableLine}}",
	   "BarCodeNumber": "{{.barcodeNumber}}"
    {{else}}
     "Errors": [
		{
			"Code": "{{.returnCode}}",
			"Message": "{{.returnMessage}}"
		}
        ]
    {{end}}
}
`

func getAPIResponseBradescoNetEmpresa() string {
	return apiResponse
}
