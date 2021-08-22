package model

import (
	"time"
)

type status string
const (
	ACTIVE	status="ACTIVE"
    INACTIVE	status="INACTIVE"
    EXPIRED	status="EXPIRED"
)

type PaymentCode struct {
	Id	string `json:"id"`
	PaymentCode	string `json:"payment_code"`
	Name	string `json:"name" validate:"required"`
	Status	*status `json:"status"`
	ExpirationDate	string `json:"expiration_date"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
}