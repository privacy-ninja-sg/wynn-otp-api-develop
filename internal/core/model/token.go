package model

type TokenExchangeReq struct {
	IdToken  string `json:"id_token" validate:"required"`
}
