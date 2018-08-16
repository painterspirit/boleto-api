package robot

import (
	"io/ioutil"
	"os"
	"time"

	"encoding/json"

	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/db"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/models"
)

type list []string

//GoRobots inicia os robôs da aplicação
func GoRobots() {
	go robotMongo(config.Get().BoletoJSONFileStore)
}

func robotMongo(path string) {
	for {
		getFiles(path).each(func(n string) {
			fileName := path + "/" + n
			data, err := ioutil.ReadFile(fileName)
			if err == nil {
				boleto := models.BoletoView{}
				errJSON := json.Unmarshal(data, &boleto)
				if errJSON == nil {
					repo, errBD := db.CreateMongo()
					if errBD == nil && repo != nil {
						errRepo := repo.SaveBoleto(boleto)
						if errRepo == nil {
							os.Remove(fileName)
						} else {
							checkError(errRepo)
						}
					} else {
						checkError(errBD)
					}
				} else {
					checkError(errJSON)
				}
			} else {
				checkError(err)
			}
		})
		time.Sleep(3600 * time.Second)
	}
}
func checkError(e error) {
	l := log.CreateLog()
	l.Warn(e, "Não foi possível recuperar arquivos de boleto")
}

func (l list) each(callback func(string)) {
	for _, name := range l {
		callback(name)
	}
}

func getFiles(path string) list {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		errCreate := os.MkdirAll(path, 0777)
		if errCreate != nil {
			return list(make([]string, 0, 0))
		}
	}
	names := make([]string, 0, len(files))
	for _, f := range files {
		names = append(names, f.Name())
	}
	return list(names)
}
