package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/models"
)

//Repository Classe de Conexão com o Banco REDIS
type Repository struct {
	conn redis.Conn
}

func (r *Repository) openConnection() error {
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

func (r *Repository) closeConnection() {
	r.conn.Close()
}

//SetBoletoHTML Grava um boleto em formato Html no Redis
func (r *Repository) SetBoletoHTML(b string, mID string) error {
	err := r.openConnection()

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%s", mID, "HTML")
	ins, _ := r.conn.Do("SETEX", key, config.Get().RedisExpirationTime, b)
	r.closeConnection()

	res := fmt.Sprintf("%s", ins)

	if res != "OK" {
		return fmt.Errorf("Could not record HTML in Redis Database: %s", res)
	}

	return nil
}

// SetBoletoJSON Grava um boleto em formato JSON no Redis
func (r *Repository) SetBoletoJSON(m *models.BoletoView) error {
	err := r.openConnection()

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s:%s", m.ID, "JSON")
	ret, _ := r.conn.Do("SETEX", key, config.Get().RedisExpirationTime, m)
	r.closeConnection()

	res := fmt.Sprintf("%s", ret)

	if res != "OK" {
		return fmt.Errorf("Could not record JSON in Redis Database: %s", res)
	}

	return nil
}

//GetMerchantTokenByKey Recupera um token da base de dados a partir do MerchantKey de um lojista
// func (r *Repository) GetMerchantTokenByKey(k string) (string, error) {
// 	err := r.openConnection()

// 	if err != nil {
// 		return "", err
// 	}

// 	key := fmt.Sprintf("%s:%s", config.Get().Application, k)
// 	ret, _ := r.conn.Do("GET", key)
// 	r.closeConnection()

// 	if ret == nil {
// 		return "", nil
// 	}

// 	return fmt.Sprintf("%s", ret), nil
// }
