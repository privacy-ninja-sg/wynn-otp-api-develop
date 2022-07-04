package model

type ReqOTPBody struct {
	PhoneNumber string `json:"phone_number"`
}

type ReqVerifyOTPBody struct {
	Token string `json:"token"`
	Pin   string `json:"pin"`
}

type RespRequestOTP struct {
	Token string `json:"token"`
	Ref   string `json:"ref"`
}
