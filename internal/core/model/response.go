package model

type DefaultResponse struct {
	Status string      `json:"s"`
	Code   int         `json:"code"`
	ErrMsg string      `json:"err_msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type TokenExchangeResponse struct {
	AccessToken string `json:"access_token"`
	Expire      int64  `json:"expire"`
}
