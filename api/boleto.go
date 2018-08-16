package api

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"strings"

	"encoding/json"
	"os"

	"io/ioutil"

	"github.com/mundipagg/boleto-api/bank"
	"github.com/mundipagg/boleto-api/boleto"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/db"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/models"
)

//Regista um boleto em um determinado banco
func registerBoleto(c *gin.Context) {

	if _, hasErr := c.Get("error"); hasErr {
		return
	}
	_boleto, _ := c.Get("boleto")
	_bank, _ := c.Get("bank")
	bol := _boleto.(models.BoletoRequest)
	bank := _bank.(bank.Bank)
	lg := bank.Log()
	lg.Operation = "RegisterBoleto"
	lg.NossoNumero = bol.Title.OurNumber
	lg.Recipient = bol.Recipient.Name
	lg.RequestKey = bol.RequestKey
	lg.BankName = bank.GetBankNameIntegration()

	resp, errR := bank.ProcessBoleto(&bol)
	if checkError(c, errR, lg) {
		return
	}

	st := http.StatusOK
	if len(resp.Errors) > 0 {

		if resp.StatusCode > 0 {
			st = resp.StatusCode
		} else {
			st = http.StatusBadRequest
		}

	} else {

		mongo, err := db.CreateMongo()
		if checkError(c, err, lg) {
			return
		}

		boView := models.NewBoletoView(bol, resp, bank.GetBankNameIntegration())
		mID, _ := boView.ID.MarshalText()
		resp.ID = string(mID)

		bhtml, err := boleto.HTML(boView, "html")
		redis := db.CreateRedis()
		err = redis.SetBoletoHTML(bhtml, resp.ID)
		if checkError(c, err, lg) {
			return
		}

		resp.Links = boView.Links
		_ = mongo.SaveBoleto(boView)
		// if errMongo != nil {

		b, err := convertToByte(boView)
		if err != nil {
			return
		}

		err = redis.SetBoletoJSON(b, resp.ID)
		if checkError(c, err, lg) {
			//LOGAR
			//MELHORAR
			return
		}

		// }
	}
	c.JSON(st, resp)
	c.Set("boletoResponse", resp)
}

// func saveBoletoJSONFile(boView models.BoletoView, lg *log.Log, err error) {
// 	lg.Warn(err.Error(), "Boleto cannot be saved at Database")
// 	fd, errOpen := os.Create(config.Get().BoletoJSONFileStore + "/boleto_" + boView.UID + ".json")
// 	if errOpen != nil {
// 		lg.Fatal(boView, "[BOLETO_ONLINE_CONTINGENCIA]"+errOpen.Error())
// 	}
// 	data, _ := json.Marshal(boView)
// 	_, errW := fd.Write(data)
// 	if errW != nil {
// 		lg.Fatal(boView, "[BOLETO_ONLINE_CONTINGENCIA]"+errW.Error())
// 	}
// 	fd.Close()
// }

func getBoleto(c *gin.Context) {
	c.Status(200)

	id := c.Query("id")
	format := c.Query("fmt")
	mongo, errCon := db.CreateMongo()
	if checkError(c, errCon, log.CreateLog()) {
		return
	}
	_boleto, err := mongo.GetBoletoByID(id)
	if err != nil {
		uid := id
		fd, err := os.Open(config.Get().BoletoJSONFileStore + "/boleto_" + uid + ".json")
		if err != nil {
			checkError(c, models.NewHTTPNotFound("Boleto não encontrado na base de dados", "MP404"), log.CreateLog())
			return
		}
		data, errR := ioutil.ReadAll(fd)
		if errR != nil {
			checkError(c, models.NewHTTPNotFound("Boleto não encontrado na base de dados", "MP404"), log.CreateLog())
			return
		}
		json.Unmarshal(data, &_boleto)
		fd.Close()
	}

	s, err := boleto.HTML(_boleto, format)
	if checkError(c, err, log.CreateLog()) {
		return
	}
	if format == "html" {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.WriteString(s)
	} else {
		c.Header("Content-Type", "application/pdf")

		buf, err := toPdf(s)

		if err != nil {
			c.Header("Content-Type", "application/json")
			checkError(c, models.NewInternalServerError(err.Error(), "internal error"), log.CreateLog())
			c.Abort()
		} else {
			c.Writer.Write(buf)
		}

	}

}

func toPdf(page string) ([]byte, error) {
	url := config.Get().PdfAPIURL
	payload := strings.NewReader(page)
	if req, err := http.NewRequest("POST", url, payload); err != nil {
		return nil, err
	} else if res, err := http.DefaultClient.Do(req); err != nil {
		return nil, err
	} else {
		defer res.Body.Close()
		return ioutil.ReadAll(res.Body)
	}
}

func getBoletoByID(c *gin.Context) {
	id := c.Param("id")
	mongo, errDb := db.CreateMongo()
	if errDb != nil {
		checkError(c, models.NewInternalServerError("MP500", "Erro interno"), log.CreateLog())
	}
	boleto, err := mongo.GetBoletoByID(id)
	if err != nil {
		checkError(c, models.NewHTTPNotFound("MP404", "Boleto não encontrado"), nil)
		return
	}
	c.JSON(http.StatusOK, boleto)
}

//convertToByte converte um model BoletoView para um JSON
func convertToByte(m models.BoletoView) (*bytes.Reader, error) {
	json, err := json.Marshal(m)
	if err != nil {
		err = errors.New("Failed to convert a json of model")
	}
	return bytes.NewReader([]byte(json)), err
}
