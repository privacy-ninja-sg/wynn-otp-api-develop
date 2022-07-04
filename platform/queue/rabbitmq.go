package queue

import (
	"github.com/streadway/amqp"
	"wynn-otp-api/pkg/configs"
	"wynn-otp-api/pkg/utils"
)

func RabbitmqConnection() *amqp.Connection {
	url := configs.RabbitmqConfig()
	conn, err := amqp.Dial(url)
	utils.FailOnErr(err)
	return conn
}
