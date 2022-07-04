package service

import (
	"github.com/parnurzeal/gorequest"
	"wynn-otp-api/internal/core/model"
	"wynn-otp-api/internal/core/service"
)

const HOST_NAME = "https://apis.shortms.com"

type shortSms struct {
	appKey    string
	appSecret string
}

func NewShortSms(appKey, appSecret string) service.ShortSMS {
	return &shortSms{appKey: appKey, appSecret: appSecret}
}

func (s shortSms) RequestOTP(phoneNumber string) (model.RequestOTPResponseBody, []error) {
	var respBody model.RequestOTPResponseBody
	reqBody := model.RequestOTPBody{
		AppKey:    s.appKey,
		SecretKey: s.appSecret,
		Msisdn:    phoneNumber,
	}
	_, _, errs := gorequest.New().
		Post(HOST_NAME + "/api/otp/request").
		SendStruct(reqBody).
		EndStruct(&respBody)
	return respBody, errs
}

func (s shortSms) VerifyOTP(token, pin string) (model.VerifyOTPResponseBody, []error) {
	var respBody model.VerifyOTPResponseBody
	reqBody := model.VerifyOTPBody{
		AppKey:    s.appKey,
		SecretKey: s.appSecret,
		Token:     token,
		Pin:       pin,
	}
	_, _, errs := gorequest.New().
		Post(HOST_NAME + "/api/otp/verify").
		SendStruct(reqBody).
		EndStruct(&respBody)
	return respBody, errs
}
