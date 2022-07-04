package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"wynn-otp-api/pkg/configs"
	"wynn-otp-api/pkg/routes"
	"wynn-otp-api/pkg/utils"
)

func init() {
	_ = godotenv.Load(".env")
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	// database connection

	// ========== Initial fiber app ==========
	app := fiber.New(configs.FiberConfig())

	// init router for routing application server...
	appEnv := os.Getenv("APP_ENV")
	router := routes.NewRouter(app, appEnv)
	// run on develop mode
	router.Logging()
	// run on all mode
	router.Main()

	// setup 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})

	// start application server with graceful shutdown...
	utils.StartServerWithGracefulShutdown(app)
}
