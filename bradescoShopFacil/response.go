package bradescoShopFacil

var apiResponse = `
{
	{{if or (eq .returnCode "0") (eq .returnCode "93005999")}}
       "DigitableLine": "{{fmtDigitableLine (trim .digitableLine)}}",
		"Links": [{
			"href":"{{.url}}",
			"rel": "pdf",
			"method":"GET"
		}]
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

func getAPIResponseBradescoShopFacil() string {
	return apiResponse
}
