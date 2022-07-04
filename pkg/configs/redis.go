package configs

import "os"

func RedisConfig() (host, port, pwd string) {
	host = os.Getenv("REDIS_HOST")
	pwd = os.Getenv("REDIS_PASSWORD")
	port = os.Getenv("REDIS_PORT")

	return host, port, pwd
}
