package utils

import (
	"LibraryManagementSys/constant"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// BcryptHash 密码加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// PasswordDecode
// 自定义密码解密
// 这里加密从简：根据时间进行对加密的密码进行切割，然后用base64解密（安全度低），返回明文
func PasswordDecode(loginTime time.Time, pw string) (string, error) {
	length := len(pw) - 3
	h := loginTime.Hour() % 10
	if h > length {
		h %= length
	}
	m := loginTime.Minute() % 100
	if m > length {
		m %= length
	}
	s := loginTime.Second() % 10
	if s > length {
		s %= length
	}
	// 按大小顺序排列
	var temp int
	if h > m {
		temp = m
		m = h
		h = temp
	}
	if m > s {
		temp = m
		m = s
		s = temp
	}
	// 删除索引指定位置元素
	pwOne := pw[0:h]
	pwTwo := pw[h+1 : m]
	pwThree := pw[m+1 : s]
	pwFour := pw[s+1:]
	// 解析出base64加密码
	encodePw := pwOne + pwTwo + pwThree + pwFour
	// 解析base64解析码为明文
	decode, err := base64.StdEncoding.DecodeString(encodePw)
	i := len(decode)
	decode = decode[:i-2]
	if err != nil {
		return constant.EMPTY, err
	} else {
		return string(decode), nil
	}
}
