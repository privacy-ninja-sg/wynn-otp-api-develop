package handler

import "github.com/gofiber/fiber/v2"

type OTPHandler interface {
	Request(c *fiber.Ctx) error
	Verify(c *fiber.Ctx) error
}
