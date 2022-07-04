package service

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"wynn-otp-api/internal/core/model"
	"wynn-otp-api/internal/core/service"
)

const OTP_HOST_NAME = "https://portal-otp.smsmkt.com"

type otpService struct {
	appKey     string
	appSecret  string
	projectKey string
}

type reqOtpBody struct {
	ProjectKey string `json:"project_key"`
	Phone      string `json:"phone"`
}

type respOtpBody struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Result struct {
		Token   string `json:"token"`
		RefCode string    `json:"ref_code"`
	}
}

type reqOtpVerifyBody struct {
	Token string `json:"token"`
	Otp   string `json:"otp_code"`
}

type respOtpVerifyBody struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Result struct {
		Status bool `json:"status"`
	} `json:"result"`
}

func NewOTPService(appKey, appSecret, otpAppProjectKey string) service.OtpService {
	return &otpService{
		appKey:     appKey,
		appSecret:  appSecret,
		projectKey: otpAppProjectKey,
	}
}

func (o otpService) RequestOTP(phoneNumber string) (model.RequestOTPResponseBody, []error) {
	var respBody model.RequestOTPResponseBody
	var respBodyRaw respOtpBody

	reqBody := reqOtpBody{
		ProjectKey: o.projectKey,
		Phone:      phoneNumber,
	}
	_, _, errs := gorequest.New().
		Post(OTP_HOST_NAME+"/api/otp-send").
		Set("content-type", "application/json").
		Set("api_key", o.appKey).
		Set("secret_key", o.appSecret).
		SendStruct(reqBody).
		//Send(`{"project_key":` + o.projectKey + `, "phone":` + phoneNumber + `}`).
		EndStruct(&respBodyRaw)

	if errs != nil {
		logrus.Error(errs)
	}

	if respBodyRaw.Code != "000" {
		return respBody, []error{errors.New(respBodyRaw.Detail)}
	}

	respBody.Success = true
	respBody.Status = "ok"
	respBody.Ref = respBodyRaw.Result.RefCode
	respBody.Token = respBodyRaw.Result.Token

	return respBody, errs
}

func (o otpService) VerifyOTP(token, pin string) (model.VerifyOTPResponseBody, []error) {
	var respBody model.VerifyOTPResponseBody
	var respBodyRaw respOtpVerifyBody

	reqBody := reqOtpVerifyBody{
		Token: token,
		Otp:   pin,
	}
	_, _, errs := gorequest.New().
		Post(OTP_HOST_NAME+"/api/otp-validate").
		SendStruct(reqBody).
		Set("api_key", o.appKey).
		Set("secret_key", o.appSecret).
		EndStruct(&respBodyRaw)

	if errs != nil {
		logrus.Error(errs)
	}

	if respBodyRaw.Code != "000" {
		return respBody, []error{errors.New(respBodyRaw.Detail)}
	}


	if respBodyRaw.Result.Status == false {
		respBody.Status = "error"
		respBody.Success = false
		respBody.Detail = "internal server error"
	} else {
		respBody.Status = "ok"
		respBody.Success = true
		respBody.Detail = respBodyRaw.Detail
	}

	return respBody, errs
}
