package service

import "wynn-otp-api/internal/core/model"

type ShortSMS interface {
	RequestOTP(phoneNumber string) (model.RequestOTPResponseBody, []error)
	VerifyOTP(token, pin string) (model.VerifyOTPResponseBody, []error)
}
