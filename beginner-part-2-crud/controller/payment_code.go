package controller

import (
	"net/http"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator"
	"github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/model"
	"github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/storage"
)

// https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
var validate *validator.Validate

func GenerateExpirationDate() time.Time {
	newExpirationDate := time.Now()
	newExpirationDate = newExpirationDate.AddDate(50, 0, 0)
	return newExpirationDate
}

func GetPaymentCodes(c echo.Context) error {
	paymentCodes, _ := GetRepoGetPaymentCodes()
	return c.JSON(http.StatusOK, paymentCodes)
}

func GetPaymentCode(c echo.Context) error {
	id := c.Param("id")
	paymentCode, err := GetRepoGetPaymentCode(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, paymentCode)
}

func CreatePaymentCode(c echo.Context) error {
	paymentCode := new(model.PaymentCode);

	if err := c.Bind(paymentCode); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	paymentCode.ExpirationDate = GenerateExpirationDate().String()

	validate = validator.New()

	if err := validate.Struct(paymentCode); err != nil {
		message := ""

		for _, err := range err.(validator.ValidationErrors) {
			message = fmt.Sprintf("%s (%s) is %s. ", err.Namespace(), err.Type(), err.Tag())
		}

		return c.JSON(http.StatusBadRequest, message)
	}

	paymentCode, err := GetRepoCreatePaymentCode(paymentCode)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, paymentCode)
}

// Repositories
func GetRepoGetPaymentCodes() ([]model.PaymentCode, error) {
	db := storage.GetDBInstance()
	paymentCodes := []model.PaymentCode{}

	// Need to explicitly specify table name here
	// if not, by convention table name is pluralized: https://gorm.io/docs/conventions.html#Pluralized-Table-Name
	if err := db.Table("payment_code").Find(&paymentCodes).Error; err != nil {
		return nil, err
	}
	return paymentCodes, nil
}

func GetRepoGetPaymentCode(id string) (*model.PaymentCode, error) {
	db := storage.GetDBInstance()
	paymentCode := &model.PaymentCode{}

	if err := db.Table("payment_code").First(paymentCode, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return paymentCode, nil
}

func GetRepoCreatePaymentCode(paymentCode *model.PaymentCode) (*model.PaymentCode, error) {
	db := storage.GetDBInstance()

	if err := db.Table("payment_code").Create(&paymentCode).Error; err != nil {
		return nil, err
	}
	return paymentCode, nil
}