package db

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/mundipagg/boleto-api/config"
)

//Redis Classe de Conexão com o Banco REDIS
type Redis struct {
	conn redis.Conn
}

var dbId = config.Get().RedisDatabase

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
func (r *Redis) SetBoletoHTML(b string, mID string) error {
	err := r.openConnection()

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%s", "HTML", mID)
	ret, _ := r.conn.Do("SETEX", key, config.Get().RedisExpirationTime, b)
	r.closeConnection()

	res := fmt.Sprintf("%s", ret)

	if res != "OK" {
		return fmt.Errorf("Could not record HTML in Redis Database: %s", res)
	}

	return nil
}

// SetBoletoJSON Grava um boleto em formato JSON no Redis
func (r *Redis) SetBoletoJSON(b *bytes.Reader, mID string) error {
	err := r.openConnection()

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%s", "JSON", mID)
	ret, _ := r.conn.Do("SET", key, b)
	r.closeConnection()

	res := fmt.Sprintf("%s", ret)

	if res != "OK" {
		return fmt.Errorf("Could not record JSON in Redis Database: %s", res)
	}

	return nil
}

//GetBoletoHTML Recupera um boleto do tipo HTML do Redis a partir do MongoID
func (r *Redis) GetBoletoHTML(mID string) (string, error) {
	err := r.openConnection()

	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("%s:%s", mID, "HTML")
	ret, _ := r.conn.Do("GET", key)
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
