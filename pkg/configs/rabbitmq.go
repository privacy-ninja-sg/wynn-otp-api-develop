package configs

import (
	"fmt"
	"os"
)

func RabbitmqConfig() string {
	host := os.Getenv("RABBITMQ_HOST")
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	port := os.Getenv("RABBITMQ_PORT")

	return fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
}
