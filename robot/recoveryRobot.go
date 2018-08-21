package robot

import (
	"github.com/jasonlvhit/gocron"
	"github.com/mundipagg/boleto-api/db"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/util"
)

//RecoveryRobot robô que faz a resiliência de boletos
func RecoveryRobot(ex string) {

	if ex == "true" {
		go func() {
			// e, _ := strconv.ParseUint(config.Get().RecoveryRobotExecutionInMinutes, 10, 64)
			gocron.Every(1).Minute().Do(executionTask)
			<-gocron.Start()
		}()
	}

}

func executionTask() {

	lg := log.CreateLog()
	lg.Operation = "RecoveryRobot"

	lg.InitRobot()

	redis := db.CreateRedis()
	keys, _ := redis.GetAllJSON()

	mongo, errMongo := db.CreateMongo(lg)
	if util.CheckErrorRobot(errMongo) == false {

		for _, key := range keys {
			bol, errRedis := redis.GetBoletoJSONByKey(string(key))
			if util.CheckErrorRobot(errRedis) == false {
				err := mongo.SaveBoleto(bol)

				if util.CheckErrorRobot(err) == false {
					lg.ResumeRobot(string(key))
					redis.DeleteBoletoJSONByKey(string(key))
				}
			}
		}
	}

	lg.EndRobot()
}
