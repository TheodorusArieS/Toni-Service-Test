package utilities

import (
	"fmt"
)

func HashThePassword(password string) string{
	newPassword := fmt.Sprintf("%s123",password)
	fmt.Println(newPassword)
	return newPassword
}