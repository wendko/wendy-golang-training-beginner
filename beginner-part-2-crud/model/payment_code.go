package model

import (
	"time"
}

type PaymentCode struct {
	Id	string `json:"id"`
	PaymentCode	string `json:"payment_code"`
	Name	string `json:"name"`
	Status	string `json:"status"`
	ExpirationDate	string `json:"expiration_date"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
}