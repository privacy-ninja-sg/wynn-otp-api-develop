package validations

import (
	"wynn-otp-api/internal/core/model"
	"wynn-otp-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateBillingValidator(c *fiber.Ctx) error {
	var reqBody model.CreateBillingReq
	err := c.BodyParser(&reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: http.StatusText(http.StatusBadGateway),
		})
	}

	validate := utils.NewValidator()
	err = validate.Struct(reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: "validate invalid!",
			Data:   utils.ValidatorErrors(err),
		})
	}

	// set request body to locals func
	c.Locals("reqBody", reqBody)

	return c.Next()
}

func UpdateBillingStatusValidator(c *fiber.Ctx) error {
	var reqBody model.UpdateBillingReq
	err := c.BodyParser(&reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: http.StatusText(http.StatusBadGateway),
		})
	}

	validate := utils.NewValidator()
	err = validate.Struct(reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: "validate invalid!",
			Data:   utils.ValidatorErrors(err),
		})
	}

	// set request body to locals func
	c.Locals("reqBody", reqBody)

	return c.Next()
}

func GetBillings(c *fiber.Ctx) error {
	var reqBody model.GetBillingReq
	err := c.QueryParser(&reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: http.StatusText(http.StatusBadGateway),
		})
	}

	validate := utils.NewValidator()
	err = validate.Struct(reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: "validate invalid!",
			Data:   utils.ValidatorErrors(err),
		})
	}

	// set request body to locals func
	c.Locals("reqBody", reqBody)

	return c.Next()
}
