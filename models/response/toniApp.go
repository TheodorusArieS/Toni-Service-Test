package response

import (
	"time"
)


type UserShopDetail struct {
	UserId         int    `json:"id,omitempty"`
	ShopOwnerName  string `json:"shop_owner_name"`
	ShopName       string `json:"shop_name"`
	ShopAddress    string `json:"shop_address"`
	ProfilePicture string `json:"profile_picture"`
	PostalCode     int    `json:"postal_code"`
}

type User struct {
	ID int `json:"id"`
	StatusID int `json:"status_id,omitempty"`
	Phone     string `json:"phone"`
	PinNumber string `json:"pin_number,omitempty"`
	UserShopDetail UserShopDetail `json:"user_shop_detail,omitempty"`
	CreatedBy int        `json:"created_by,omitempty"`
	UpdatedBy int        `json:"updated_by,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Token     string     `json:"token,omitempty"`
}

type RestResponse struct {
	Data interface{}
	Message string `json:"message"`
	Status int `json:"status"`
}

type LoginData struct {
	ID int `json:"id"`
	Phone string `json:"phone"`
	PinNumber string `json:"pin_number"`
}