package boleto

import (
	"bytes"
	"encoding/base64"
	"errors"
	"html/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/twooffive"

	"image/png"

	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/tmpl"
)

const templateBoleto = `
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <style>

        @media print
        {    
            .no-print, .no-print *
            {
                display: none !important;
            }
        }

        body {
            font-family: "Arial";
    		background-color: #fff;
            font-size:0.7em;
        }
        .left {
    		margin: auto;		
    		width: 216mm;
    	}
        .document {
            margin: auto auto;
            width: 216mm;
            height: 87mm;
        }

        .headerBtn {
            margin: auto auto;
            width: 216mm;
            background-color: #fff;
        }

        table {
            width: 100%;
            position: relative;
            border-collapse: collapse;
        }

        .boletoNumber {
            width: 66%;
            font-weight: bold;
            font-size:0.9em;
        }

        .center {
            text-align: center;
        }

        .right {
            text-align: right;
            right: 20px;
        }

        td {
            position: relative;
        }

        .title {
            position: absolute;
            left: 0px;
            top: 0px;
            font-size:0.65em;
            font-weight: bold;
        }

        .text {
             font-size:0.7em;
        }

        p.content {
            padding: 0px;
            width: 100%;
            margin: 0px;
            font-size:0.7em;
        }

        .sideBorders {
            border-left: 1px solid black;
            border-right: 1px solid black;
        }

        hr {
            size: 1;
            border: 1px dashed;
    		width: 216mm;
    		margin-top: 9mm;
        	margin-bottom: 9mm;
        }

        br {
            content: " ";
            display: block;
            margin: 12px 0;
            line-height: 12px;
        }

        .print {
            /* TODO(dbeam): reconcile this with overlay.css' .default-button. */
            background-color: rgb(77, 144, 254);
            background-image: linear-gradient(to bottom, rgb(77, 144, 254), rgb(71, 135, 237));
            border: 1px solid rgb(48, 121, 237);
            color: #fff;
            text-shadow: 0 1px rgba(0, 0, 0, 0.1);
        }

        .btnDefault {
            font-kerning: none;
            font-weight: bold;
        }

        .btnDefault:not(:focus):not(:disabled) {
            border-color: #808080;
        }

        button {
            border: 1px;
            padding: 5px;
            line-height: 20px;
        }

        span.iconFont {
            font-size: 20px;
        }

        span.align {
            display: inline-block;
            vertical-align: middle;
        }
    </style>
    <link rel="stylesheet" href="https://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css">
</head>

<body>
    {{if eq .Format "html"}}	
	<br/>
    <div class="headerBtn">
        <div style="text-align:right;">
            <button class="no-print btnDefault print" onclick="window.print()">
                <span class="align iconFont ion-printer"></span>
                <span class="align">&nbspImprimir</span>
            </button>
            <button class="no-print btnDefault print" onclick="window.location='./boleto?fmt=pdf&id={{.ID}}'">
                <span class="align iconFont ion-document-text"></span>
                <span class="align">&nbspGerar PDF</span>
            </button>
            <!--<button class="no-print btnDefault print" onclick="window.location='./boleto/www.google.com'">
                <span class="align iconFont ion-image"></span>
                <span class="align">&nbspSalvar como Imagem</span>
            </button>-->
        </div>
    </div>
    <br/>
    {{end}}
    {{template "boletoForm" .}}

	<hr/>
	{{template "boletoForm" .}}
	<div class="left">
		<img style="margin-left:5mm;" id="barcode_{{printIfproduction .Barcode}}" src="data:image/png;base64,{{.Barcode64}}" alt="">
		<br/>
		</div>
    </div>
</body>

</html>
`

const boletoForm = `
{{define "boletoForm"}}
<div class="document">
        <table cellspacing="0" cellpadding="0">
            <tr class="topLine">
                <td class="bankLogo">
                    {{.BankLogo}}					
                </td>
                <td class="sideBorders center"><span style="font-weight:bold;font-size:0.9em;">{{.BankNumber}}</span></td>
                <td class="boletoNumber center"><span>{{.DigitableLine}}</span></td>
            </tr>
        </table>
        <table cellspacing="0" cellpadding="0" border="1">
            <tr>
                <td width="70%" colspan="6">
                    <span class="title">Local de Pagamento</span>
                    <br/>
                    <span class="text">ATÉ O VENCIMENTO EM QUALQUER BANCO OU CORRESPONDENTE NÃO BANCÁRIO</span>
                </td>
                <td width="30%">
                    <span class="title">Data de Vencimento</span>
                    <br/>
                    <br/>
                    <p class="content right text" style="font-weight:bold;" id="expire_date">{{.Boleto.Title.ExpireDateTime | brdate}}</p>
                </td>
            </tr>
            <tr>
                <td width="70%" colspan="6">
                    <span class="title">Nome do Beneficiário / CNPJ / CPF / Endereço:</span>
                    <br/>
                    <table border="0" style="border:none">
                        <tr>
                            <td width="60%"><span class="text" id="recipient_name">{{.Boleto.Recipient.Name}}</span></td>
                            <td><span class="text" id="recipient_document"><b>{{.Boleto.Recipient.Document.Type}}</b> {{fmtDoc .Boleto.Recipient.Document}}</span></td>
                        </tr>
                    </table>
                    <br/>
                    <span class="text" id="recipient_address">{{.Boleto.Recipient.Address.Street}}, 
                    {{.Boleto.Recipient.Address.Number}} - 
                    {{.Boleto.Recipient.Address.District}}, 
                    {{.Boleto.Recipient.Address.StateCode}} - 
                    {{.Boleto.Recipient.Address.ZipCode}}</span>
                </td>
                <td width="30%">
                    <span class="title">Agência/Código Beneficiário</span>
                    <br/>
                    <br/>
                    <p class="content right" id="agreement_agency_account">{{.Boleto.Agreement.Agency}}/{{.Boleto.Agreement.Account}}</p>
                </td>
            </tr>

            <tr>
                <td width="15%">
                    <span class="title">Data do Documento</span>
                    <br/>
                    <p class="content center" id="create_date">{{.Boleto.Title.CreateDate | brdate}}</p>
                </td>
                <td width="17%" colspan="2">
                    <span class="title">Num. do Documento</span>
                    <br/>
                    <p class="content center" id="boleto_document_number">{{.Boleto.Title.DocumentNumber}}</p>
                </td>
                <td width="10%">
                    <span class="title">Espécie doc</span>
                    <br/>
                    <p class="content center" id="configbank_especie_doc">{{.ConfigBank.EspecieDoc}}</p>
                </td>
                <td width="8%">
                    <span class="title">Aceite</span>
                    <br/>
                    <p class="content center" id="configbank_aceite" >{{.ConfigBank.Aceite}}</p>
                </td>
                <td>
                    <span class="title">Data Processamento</span>
                    <br/>
                    <p class="content center" id="process_date">{{.Boleto.Title.CreateDate | brdate}}</p>
                </td>
                <td width="30%">
                    <span class="title">Carteira/Nosso Número</span>
                    <br/>
                    <br/>
                    <p class="content right" id="ournumber">{{.Boleto.Agreement.Wallet}}/{{.Boleto.Title.OurNumber}}</p>
                </td>
            </tr>

            <tr>
                <td width="15%">
                    <span class="title">Uso do Banco</span>
                    <br/>
                    <p class="content center">&nbsp;</p>
                </td>
                <td width="10%">
                    <span class="title">Carteira</span>
                    <br/>
                    <p class="content center" id="wallet">{{.Boleto.Agreement.Wallet}}</p>
                </td>
                <td width="10%">
                    <span class="title">Espécie</span>
                    <br/>
                    <p class="content center">R$</p>
                </td>
                <td width="8%" colspan="2">
                    <span class="title">Quantidade</span>
                    <br/>
                    <p class="content center" id="configbank_quantidade">{{.ConfigBank.Quantidade}}</p>
                </td>
                <td>
                    <span class="title">Valor</span>
                    <br/>
                    <p class="content center" id="configbank_valorCotacao" >{{.ConfigBank.ValorCotacao}}</p>
                </td>
                <td width="30%">
                    <span class="title">(=) Valor do Documento</span>
                    <br/>
                    <br/>
                    <p class="content right" id="amount_in_cents" >{{fmtNumber .Boleto.Title.AmountInCents}}</p>
                </td>
            </tr>
            <tr>
                <td colspan="6" rowspan="4">
                    <span class="title">Instruções de responsabilidade do BENEFICIÁRIO. Qualquer dúvida sobre este boleto contate o beneficiário.</span>
                    <p class="content" id="instructions">{{unescapeHtmlString .Boleto.Title.Instructions}}</p>
                </td>
            </tr>
            <tr>
                <td>
                    <span class="title">(-) Descontos/Abatimento</span>
                    <br/>
                    <p class="content right">&nbsp;</p>
                </td>
            </tr>
            <tr>
                <td>
                    <span class="title">(+) Juros/Multa</span>
                    <br/>
                    <p class="content right">&nbsp;</p>
                </td>
            </tr>
            <tr>
                <td>
                    <span class="title">(=) Valor Pago</span>
                    <br/>
                    <p class="content right">&nbsp;</p>
                </td>
            </tr>
            <tr>
                <td colspan="7">
                    <table border="0" style="border:none">
                        <tr>
                            <td width="60%"><span class="text" id="buyer_name"><b>Nome do Pagador: </b>&nbsp;{{.Boleto.Buyer.Name}}</span></td>
                            <td><span class="text" id="buyer_document"><b>CNPJ/CPF: </b>&nbsp;{{fmtDoc .Boleto.Buyer.Document}}</span></td>
                        </tr>
                        <tr>
                            <td><span class="text" id="buyer_address"><b>Endereço: </b>&nbsp;{{.Boleto.Buyer.Address.Street}}&nbsp;{{.Boleto.Buyer.Address.Number}}, {{.Boleto.Buyer.Address.District}} - {{.Boleto.Buyer.Address.City}}, {{.Boleto.Buyer.Address.StateCode}} - {{.Boleto.Buyer.Address.ZipCode}}</span></td>
                            <td>&nbsp;</td>
                        </tr>
                        <tr>
                            <td><span class="text"><b>Sacador/Avalista: </b> &nbsp;</span></td>
                            <td><span class="text"><b>CNPJ/CPF: </b> &nbsp;</span></td>
                        </tr>
                    </table>

                </td>

            </tr>
        </table>
		<br/>
    </div>

	{{end}}
`

//HTML renderiza HTML do boleto
func HTML(boleto models.BoletoView, format string) (string, error) {
	if boleto.Barcode == "" {
		return "", errors.New("boleto not found")
	}
	b := tmpl.New()
	boleto.BankLogo = template.HTML(boleto.ConfigBank.Logo)
	boleto.Format = format
	bcode, _ := twooffive.Encode(boleto.Barcode, true)
	orgBounds := bcode.Bounds()
	orgWidth := orgBounds.Max.X - orgBounds.Min.X
	img, _ := barcode.Scale(bcode, orgWidth, 50)
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	boleto.Barcode64 = base64.StdEncoding.EncodeToString(buf.Bytes())
	s, err := b.From(boleto).To(templateBoleto).Transform(boletoForm)
	if err != nil {
		return "", err
	}
	return s, nil
}
