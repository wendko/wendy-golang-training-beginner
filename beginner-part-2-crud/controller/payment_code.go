package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/model"
	"github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/storage"
)

func GetPaymentCodes(c echo.Context) error {
	paymentCodes, _ := GetRepoGetPaymentCodes()
	return c.JSON(http.StatusOK, paymentCodes)
}

// func FindOnePaymentCode(c echo.Context) error {
// 	paymentCode, _ := GetRepoFindPaymentCode()
// 	return c.JSON(http.StatusOK, paymentCode)
// }

// func CreatePaymentCode(c echo.Context) error {
// 	paymentCodes, _ := GetRepoGetPaymentCodes()
// 	return c.JSON(http.StatusOK, paymentCodes)
// }

func GetRepoGetPaymentCodes() ([]model.PaymentCode, error) {
	db := storage.GetDBInstance()
	paymentCode := []model.PaymentCode{}

	// Need to explicitly specify table name here
	// if not, by convention table name is pluralized: https://gorm.io/docs/conventions.html#Pluralized-Table-Name
	if err := db.Table("payment_code").Find(&paymentCode).Error; err != nil {
		return nil, err
	}
	return paymentCode, nil
}

// func GetRepoFindPaymentCode(id string) ([]model.PaymentCode, error) {
// 	db := storage.GetDBInstance()
// 	paymentCodes := []model.PaymentCode{}

// 	if err := db.Find(&paymentCodes).Error; err != nil {
// 		return nil, err
// 	}
// 	return paymentCodes, nil
// }