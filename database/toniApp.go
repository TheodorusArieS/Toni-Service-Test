package database

import (
	"fmt"
	"toni-service-test/database/schema"
	"toni-service-test/models/response"

	// "toni-service-test/database"
	"toni-service-test/query"
)

type ToniApp struct {
}

func (ta *ToniApp) GetUserWhere(query string, args ...interface{}) (response.User, error) {
	var result response.User
	err := Database.Table(schema.User{}.TableName()).Select(`id,status_id,phone,pin_number,created_at,updated_at,created_by,updated_by`).
		Where(query, args...).
		Scan(&result).Error
	return result, err
}

func (ta *ToniApp) CreateOTP(otpVerification *schema.OTPVerification) error {
	otp := otpVerification.OTP
	phone := otpVerification.Phone
	typeID := otpVerification.TypeId
	stmt, err := Client.Prepare(query.QueryCreateOTP)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(phone, otp, typeID)
	if err != nil {
		return err
	} else {
		fmt.Println("Berhasil Create OTP")
		return nil
	}
}

func (ta *ToniApp) ValidateOTP(otpVerification *schema.OTPVerification) bool {

	phone := otpVerification.Phone
	otp := otpVerification.OTP
	typeID := otpVerification.TypeId

	var result int
	stmt, err := Client.Prepare(query.QueryValidateOTP)
	if err != nil {
		fmt.Println("ERROR WHEN PREPARING QUERY")
		return false
	}
	data, err := stmt.Query(phone, otp, typeID)
	if err != nil {
		fmt.Println("ERROR WHEN TAKING DATA FROM DATABASE")
		return false
	}

	for data.Next() {
		var res schema.OTPVerification
		if err := data.Scan(&res.Id); err != nil {
			fmt.Println("ERROR WHEN SCANNING 123")
			return false
		}
		result = res.Id
	}
	fmt.Println(result)

	if result != 0 {
		fmt.Println("berhasil")
		stmt, err = Client.Prepare(query.QueryDeleteValidateOTP)
		if err != nil {
			fmt.Println("ERROR PREPARING DELETE QUERY")
		}
		_, err = stmt.Exec(result)
		if err != nil {
			fmt.Println("ERROR WHEN EXEC DATA")
		}
		return true
	} else {
		fmt.Println("ada di sini")
		return false
	}

}

func (ta *ToniApp) RegisterUser(user *schema.User) error {
	phone := user.Phone
	pinNumber := user.PinNumber
	StatusID := user.StatusID

	stmt, err := Client.Prepare(query.QueryCreateNewUser)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(phone, pinNumber, StatusID)
	if err != nil {
		return err
	}
	return nil
}

func (ta *ToniApp) Login(user *schema.User) (*response.User, error) {
	stmt, err := Client.Prepare(query.QueryLogin)
	// result := &response.RestResponse{
	// 	Data:    nil,
	// 	Status:  200,
	// 	Message: "Semoga sukses",
	// }
	res :=& response.User{}
	if err != nil {
		fmt.Println("ERROR IN PREPARING QUERY")
		return res, err
	}
	dbResponse := stmt.QueryRow(user.Phone)

	// if err != nil {
	// 	fmt.Println("ERROR IN GETTING DATA FROM DATABASE")
	// 	return result, err
	// }


	// for dbResponse.Next() {

	if err := dbResponse.Scan(&res.ID,&res.Phone, &res.PinNumber); err != nil {
		fmt.Printf("Error in scan %#v", err.Error())
		// if getErr.Error() == "sql: no rows in result set"{
		// 	return result, err
		// }
		if err.Error() == "sql: no rows in result set"{
			fmt.Println("ada di dalam sini")
		}
		
	}

	// }
	fmt.Println("ERROR IN LAST RETURN")
	fmt.Println(res)
	return res, err

}
