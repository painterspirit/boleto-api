package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"
)

//Redis Classe de Conexão com o Banco REDIS
type Redis struct {
	conn redis.Conn
}

//CreateRedis Cria instancia do Struct Redis
func CreateRedis() *Redis {
	return new(Redis)
}

func (r *Redis) openConnection() error {
	dbID, _ := strconv.Atoi(config.Get().RedisDatabase)
	o := redis.DialDatabase(dbID)
	ps := redis.DialPassword(config.Get().RedisPassword)
	tOut := redis.DialConnectTimeout(15 * time.Second)

	c, err := redis.Dial("tcp", config.Get().RedisURL, o, ps, tOut)
	if err != nil {
		return err
	}

	r.conn = c
	return nil
}

func (r *Redis) closeConnection() {
	r.conn.Close()
}

//SetBoletoHTML Grava um boleto em formato Html no Redis
func (r *Redis) SetBoletoHTML(b string, mID string, lg *log.Log) {
	err := r.openConnection()

	if err != nil {
		lg.Warn(err, fmt.Sprintf("OpenConnection - Could not connection to Redis Database"))
	} else {

		key := fmt.Sprintf("%s:%s", "HTML", mID)
		ret, err := r.conn.Do("SETEX", key, config.Get().RedisExpirationTime, b)
		r.closeConnection()

		res := fmt.Sprintf("%s", ret)

		if res != "OK" {
			lg.Warn(err, fmt.Sprintf("SetBoletoHTML - Could not record HTML in Redis Database: %s", err.Error()))
		}
	}

}

//SetBoletoJSON Grava um boleto em formato JSON no Redis
func (r *Redis) SetBoletoJSON(b, mID string) error {
	err := r.openConnection()

	if err != nil {
		return err
	} else {
		key := fmt.Sprintf("%s:%s", "JSON", mID)
		ret, err := r.conn.Do("SET", key, b)
		r.closeConnection()

		res := fmt.Sprintf("%s", ret)

		if res != "OK" {
			return err
		}
	}
	return nil
}

//GetBoletoHTMLByID busca um boleto pelo ID que vem na URL
func (r *Redis) GetBoletoHTMLByID(id string) (string, error) {
	err := r.openConnection()

	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("%s:%s", "HTML", id)
	ret, err := r.conn.Do("GET", key)
	r.closeConnection()

	if ret == nil {
		return "", nil
	}

	return fmt.Sprintf("%s", ret), nil
}

//GetBoletoJSON Reupero um boleto do tipo JSON do Redis a partir do MongoID
// func (r *Redis) GetBoletoJSON(mID string) (m *models.BoletoView, error) {
// 	err := r.openConnection()

// 	if err != nil {
// 		return "", err
// 	}

// 	key := fmt.Sprintf("%s:%s", mID, "JSON")
// 	ret, _ := r.conn.Do("GET", key)
// 	r.closeConnection()

// 	if ret == nil {
// 		return "", nil
// 	}

// 	return m, nil
// }
