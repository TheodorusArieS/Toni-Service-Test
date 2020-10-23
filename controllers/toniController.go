package controllers

import (
	// "fmt"
	"math/rand"
	"strings"
	"time"
	"toni-service-test/constant/status"
	types "toni-service-test/constant/type"
	"toni-service-test/database"
	"toni-service-test/database/schema"
	"toni-service-test/models/request"

	"github.com/gin-gonic/gin"
)

type ToniAppController struct {
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func removeLeadingZero(input string) string {
	s := []rune(input)
	res := delChar(s, 3)
	return string(res)
}

func rangeIn(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func (ta *ToniAppController) Register(c *gin.Context) {
	newRegisterData := &request.Register{}
	_ = c.ShouldBind(&newRegisterData)

	helper := Helper{c}

	validation := newRegisterData.ValidateRequest()

	if validation != nil {
		helper.BadRequest(validation, "Invalid Data Input")
	}

	if strings.Contains(newRegisterData.Phone, "+620") {
		newRegisterData.Phone = removeLeadingZero(newRegisterData.Phone)
	}

	toniAppModel := database.ToniApp{}

	isValid := toniAppModel.ValidateOTP(
		&schema.OTPVerification{
			Phone:  newRegisterData.Phone,
			OTP:    newRegisterData.OTP,
			TypeId: types.OTP_REGISTER,
		},
	)

	if isValid {
		helper.SuccessResponse(isValid, "OTP Ketemu")
		newUser := &schema.User{
			Phone:     newRegisterData.Phone,
			PinNumber: newRegisterData.PinNumber,
			StatusID:  status.REGISTER,
		}
		newUser.HashPin()
		toniAppModels := database.ToniApp{}
		_ = toniAppModels.RegisterUser(newUser)
		helper.SuccessResponse("Data Berhasil di tambahkan", "Berhasil")
	} else {
		helper.BadRequest(isValid, "INVALID DATA")
	}
}

func (ta *ToniAppController) CreateOTP(c *gin.Context) {
	createOTPRequest := request.CreateOTP{}
	_ = c.ShouldBind(&createOTPRequest)

	helper := Helper{c}

	validation := createOTPRequest.ValidateRequest()

	if validation != nil {
		helper.BadRequest(validation, "Invalid Validation")
	}

	if strings.Contains(createOTPRequest.Phone, "+620") {
		createOTPRequest.Phone = removeLeadingZero(createOTPRequest.Phone)
	}

	toniAppModel := database.ToniApp{}

	_, err := toniAppModel.Login(&schema.User{
		Phone: createOTPRequest.Phone,
	})

	if createOTPRequest.Type == types.OTP_REGISTER {
		if err == nil {
			helper.BadRequest(createOTPRequest.Phone, "Invalid Request")
			return
		}
	}
	if createOTPRequest.Type == types.OTP_FORGET_PASSWORD {
		if err != nil {
			helper.BadRequest(createOTPRequest.Phone, "Invalid Request")
			return
		}
	}

	newOTP := rangeIn(100000, 999999)

	err = toniAppModel.CreateOTP(&schema.OTPVerification{
		OTP:    newOTP,
		Phone:  createOTPRequest.Phone,
		TypeId: createOTPRequest.Type,
	})

	if err == nil {
		helper.SuccessResponse(newOTP, "Success Created OTP")
	} else {
		helper.BadRequest(err, "Invalid Create OTP")
	}

}

func (ta *ToniAppController) Login(c *gin.Context) {
	loginData := &request.Login{}
	_ = c.ShouldBind(&loginData)
	helper := Helper{c}
	validation := loginData.ValidateRequest()
	if validation != nil {
		helper.BadRequest(validation, "ERROR VALIDATION LOGIN")
	}

	if strings.Contains(loginData.Phone, "+620") {
		loginData.Phone = removeLeadingZero(loginData.Phone)
	}

	ToniAppModels := database.ToniApp{}

	dbResponse, err := ToniAppModels.Login(&schema.User{
		Phone: loginData.Phone,
	})
	if err != nil {
		helper.BadRequest(err, "ERROR DATA FROM DATABASE")
	} else {
		helper.SuccessResponse(dbResponse, "SUCCESS DATA FROM DATABASE")

	}
}
