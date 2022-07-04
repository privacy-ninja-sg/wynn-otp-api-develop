package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"wynn-otp-api/internal/core/handler"
	"wynn-otp-api/internal/core/model"
	"wynn-otp-api/internal/core/service"
	"wynn-otp-api/pkg/errMsg"
)

type otpHandler struct {
	otpService service.OtpService
}

func NewOTPHandler(otpServ service.OtpService) handler.OTPHandler {
	return &otpHandler{otpService: otpServ}
}

// Request : post
func (o otpHandler) Request(c *fiber.Ctx) error {
	var reqBody model.ReqOTPBody
	err := c.BodyParser(&reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: errMsg.INVALID_REQUEST,
		})
	}

	resp, errs := o.otpService.RequestOTP(reqBody.PhoneNumber)
	if errs != nil {
		return c.JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: errs[0].Error(),
		})
	}

	return c.JSON(model.DefaultResponse{
		Status: "ok",
		Code:   http.StatusOK,
		Data: model.RespRequestOTP{
			Token: resp.Token,
			Ref:   resp.Ref,
		},
	})
}

// Verify : post
func (o otpHandler) Verify(c *fiber.Ctx) error {
	var reqBody model.ReqVerifyOTPBody
	err := c.BodyParser(&reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: errMsg.INVALID_REQUEST,
		})
	}

	resp, errs := o.otpService.VerifyOTP(reqBody.Token, reqBody.Pin)
	if errs != nil {
		return c.JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: errs[0].Error(),
		})
	}

	if resp.Success == false {
		return c.Status(http.StatusBadRequest).JSON(model.DefaultResponse{
			Status: "error",
			Code:   http.StatusBadRequest,
			ErrMsg: resp.Error.Detail,
		})
	}

	return c.Status(http.StatusOK).JSON(model.DefaultResponse{
		Status: "ok",
		Code:   http.StatusOK,
		Data: fiber.Map{
			"detail": resp.Detail,
		},
	})

}
