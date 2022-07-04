package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"wynn-otp-api/app/handler"
	"wynn-otp-api/app/service"
)

type Router struct {
	app    *fiber.App
	appEnv string
}

func NewRouter(app *fiber.App, appEnv string) *Router {
	return &Router{app: app, appEnv: appEnv}
}

func (r *Router) Logging() {
	if r.appEnv == "develop" {
		r.app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
	}
}

func (r *Router) Main() {
	otpAppKey := os.Getenv("OTP_APP_KEY")
	otpAppSecret := os.Getenv("OTP_APP_SECRET")
	otpAppProjectKey := os.Getenv("OTP_APP_PROJECT_KEY")
	appKey := os.Getenv("APP_KEY")
	appSecret := os.Getenv("APP_SECRET")

	// new service
	otpService := service.NewOTPService(otpAppKey, otpAppSecret, otpAppProjectKey)

	// new handler
	hand := handler.NewOTPHandler(otpService)

	// un-auth endpoint
	r.app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"s":   "ok",
			"msg": "hello world",
		})
	})

	r.app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			appKey: appSecret,
		},
	}))

	r.app.Post("/api/request", hand.Request)
	r.app.Post("/api/verify", hand.Verify)
}
