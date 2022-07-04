package model

type RequestOTPResponseBody struct { // shortSMS
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Token   string `json:"token"`
	Ref     string `json:"ref"`
}

type VerifyOTPResponseBody struct { // shortSMS
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Detail  string `json:"detail"`
	Error   struct {
		Status string `json:"status"`
		Detail string `json:"detail"`
	} `json:"error,omitempty"`
}

type RequestOTPBody struct { // shortSMS
	AppKey    string `json:"appKey"`
	SecretKey string `json:"secretKey"`
	Msisdn    string `json:"msisdn"`
}

type VerifyOTPBody struct { // shortSMS
	AppKey    string `json:"appKey"`
	SecretKey string `json:"secretKey"`
	Token     string `json:"token"`
	Pin       string `json:"pin"`
}
