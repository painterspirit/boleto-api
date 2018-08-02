package mock

import (
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func registerBoletoBradescoShopFacil(c *gin.Context) {

	const tok = `
{
    "merchant_id": "90000",
    "meio_pagamento": "800",
    "pedido": {
        "numero": "0-9_A-Z_.MAX-27-CH99",
        "valor": 15000,
        "descricao": "Descritivo do pedido"
    },
    "boleto": {
        "valor_titulo": 15000,
        "data_geracao": "2016-04-22T08:10:43",
        "data_atualizacao": null,
        "linha_digitavel": "23792372215000460151949000560000176050000013114",
        "linha_digitavel_formatada": "23792.37221  50004.601519  49000.560000  1  76050000013114",
        "token": "c3ZtRGVKRDFoUlRESmxRNnhKQnpJalFrb0VueXdVdUxnT2FVMG45cm1qMFMyRDcwRWZ0cFVBS0o0\nMFAxOHY0aTdJK3E1MXVjUVJjNEpBdUxvcE15T1E9PQ==",
        "url_acesso": "http://localhost:9080/boleto/titulo?token=c3ZtRGVKRDFoUlRESmxRNnhKQnpJalFrb0VueXdVdUxnT2FVMG45cm1qMFMyRDcwRWZ0cFVBS0o0\nMFAxOHY0aTdJK3E1MXVjUVJjNEpBdUxvcE15T1E9PQ=="
    },
    "status": {
        "codigo": 0,
        "mensagem": "OPERACAO REALIZADA COM SUCESSO",
        "detalhes": null
    }
}
`

	const tokErr = `
{
    "merchant_id": "90000",
    "meio_pagamento": "300",
    "pedido": {
        "numero": "0-9_A-Z_.MAX-27-CH99",
        "valor": 15000,
        "descricao": "Descritivo do pedido"
    },
    "boleto": null,
    "status": {
        "codigo": -518,
        "mensagem": "Erro - Mock BradescoShopFacil",
        "detalhes": "Erro - Mock BradescoShopFacil"
    }
}
`
	d, _ := ioutil.ReadAll(c.Request.Body)
	json := string(d)
	if strings.Contains(json, `valor": 200,`) {
		c.Data(200, "text/json", []byte(tok))
	} else {
		c.Data(200, "text/json", []byte(tokErr))
	}
}
