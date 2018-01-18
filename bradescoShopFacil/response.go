package bradescoShopFacil

var apiResponse = `
{
	{{if eq .returnCode "0"}}
       "DigitableLine": "{{.digitableLine}}",
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
