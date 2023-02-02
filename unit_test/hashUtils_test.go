package unit_test

import (
	"LibraryManagementSys/model/request"
	"LibraryManagementSys/utils"
	"fmt"
	"testing"
	"time"
)

// TestPasswordDecode
// 对密码解析进行单元测试
func TestPasswordDecode(t *testing.T) {
	toBeCharge := "2023-01-31 13:58:18"
	layout := "2006-01-02 15:04:05"
	location, _ := time.LoadLocation("Asia/Shanghai")
	inLocation, _ := time.ParseInLocation(layout, toBeCharge, location)
	user := request.LoginUser{LoginTime: inLocation, Password: "MTI1zNDUp2oCiA=", Username: "admin"}
	decode, _ := utils.PasswordDecode(user.LoginTime, user.Password)
	fmt.Println(decode)
}

func TestBcryptHash(t *testing.T) {
	hash := utils.BcryptHash("123456")
	fmt.Println(hash)
	check := utils.BcryptCheck("123456", hash)
	fmt.Println(check)
}

func TestBcryptCheck(t *testing.T) {
	check := utils.BcryptCheck("123456", "$2a$10$qmzRBabkFoz1d4Y3qDMu7ec3lgERA03cc2i21nYSkKLKX5JxY7tJm")
	fmt.Println(check)
}
