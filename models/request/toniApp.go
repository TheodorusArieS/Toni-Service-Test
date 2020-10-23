package request

type Register struct{
	OTP int `json:"otp"`
	Phone string `json:"phone"`
	PinNumber string `json:"pin_number"`
}

type CreateOTP struct {
	Phone string `json:"phone"`
	Type int `json:"type"`
}

type Login struct {
	Phone string `json:"phone"`
	PinNumber string `json:"pin_number"`
}
