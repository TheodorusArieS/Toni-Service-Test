package schema

type OTPVerification struct {
	Base
	Phone  string `gorm:"unique;not_null"`
	OTP    int    `gorm:"not_null"`
	TypeId int    `gorm:"not_null"`
}

func (OTPVerification) TableName() string {
	return "otp_verification"
}
func (OTPVerification) Pk() string {
	return "id"
}

func (o OTPVerification) Ref() string {
	return o.TableName() + "(" + o.Pk() + ")"
}

func (o OTPVerification) AddForeignKeys(){
	
}

func (o OTPVerification) InsertDefaults(){

}


