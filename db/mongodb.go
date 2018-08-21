package db

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoDb Struct
type MongoDb struct {
	m sync.RWMutex
}

var dbName = "Boleto"

var (
	dbSession *mgo.Session
	err       error
)

//CreateMongo cria uma nova intancia de conexão com o mongodb
func CreateMongo(l *log.Log) (*MongoDb, error) {

	if dbSession == nil {
		dbSession, err = mgo.DialWithInfo(getInfo())

		if err != nil {
			l.Warn(err, fmt.Sprintf("Error create connection mongo %s", err.Error()))
			return nil, err
		}
	}

	db := new(MongoDb)
	// if config.Get().MockMode {
	// 	dbName = "boletoapi_mock"
	// }
	return db, nil
}

func getInfo() *mgo.DialInfo {
	connMgo := strings.Split(config.Get().MongoURL, ",")
	return &mgo.DialInfo{
		Addrs:     connMgo,
		Timeout:   10 * time.Second,
		Database:  "Boleto",
		PoolLimit: 512,
		Username:  config.Get().MongoUser,
		Password:  config.Get().MongoPassword,
	}
}

//SaveBoleto salva um boleto no mongoDB
func (e *MongoDb) SaveBoleto(boleto models.BoletoView) error {

	e.m.Lock()
	defer e.m.Unlock()

	session := dbSession.Copy()

	defer session.Close()

	c := session.DB(dbName).C("boletos")
	err = c.Insert(boleto)
	return err
}

//GetBoletoByID busca um boleto pelo ID que vem na URL
func (e *MongoDb) GetBoletoByID(id string) (models.BoletoView, error) {
	e.m.Lock()
	defer e.m.Unlock()
	result := models.BoletoView{}

	session := dbSession.Copy()

	defer session.Close()

	c := session.DB(dbName).C("boletos")

	if len(id) == 24 {
		d := bson.ObjectIdHex(id)
		err = c.Find(bson.M{"_id": d}).One(&result)
	} else {
		err = c.Find(bson.M{"id": id}).One(&result)
	}

	if err != nil {
		l := log.CreateLog()
		l.Warn(err, fmt.Sprintf("GetBoletoByID %s", err.Error()))
		return models.BoletoView{}, err
	}

	return result, nil
}

func (e *MongoDb) Close() {
	fmt.Println("Close Database Connection")
}
