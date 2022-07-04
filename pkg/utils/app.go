// ./pkg/utils/start_server.go

package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(app *fiber.App) {
	appPort := os.Getenv("APP_PORT")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// ...

	if err := app.Listen(":" + appPort); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}
