package models

type PostJSON struct {
	LongURL string `json:"longURL" binding:"required"`
	Expiry  string `json:"expiry"`
}
