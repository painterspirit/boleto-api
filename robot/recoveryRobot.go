package robot

import (
	"strconv"

	"github.com/jasonlvhit/gocron"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"
)

//RecoveryRobot robô que faz a resiliência de boletos
func RecoveryRobot(ex string) {

	if ex == "true" {
		log.InitRobot()

		go func() {
			e, _ := strconv.ParseUint(config.Get().RecoveryRobotExecutionInMinutes, 10, 64)
			gocron.Every(e).Minutes().Do(executionTask)
			<-gocron.Start()
		}()

	}

}

func executionTask() {
	log.Info("deu")
}
