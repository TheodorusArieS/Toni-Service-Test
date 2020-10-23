package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v *Register)ValidateRequest() error {
	return validation.ValidateStruct(
		v,
		validation.Field(&v.OTP,validation.Required),
		validation.Field(&v.Phone, validation.Required),
		validation.Field(&v.PinNumber, validation.Required,validation.Length(6,6)),
	)
}

func (c *CreateOTP) ValidateRequest() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Phone,validation.Required),
		validation.Field(&c.Type,validation.Required),
	)
}

func (c *Login) ValidateRequest() error{
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Phone,validation.Required),
		validation.Field(&c.PinNumber,validation.Required),
	)
}