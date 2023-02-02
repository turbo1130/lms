package service

import (
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
	"LibraryManagementSys/model/request"
	systemRes "LibraryManagementSys/model/response"
	"LibraryManagementSys/utils"
	"context"
	"encoding/json"
	"errors"
	"time"
)

type UserService struct {
}

func (u *UserService) Login(user model.User) (userRes systemRes.User, err error) {
	var userSelect model.User
	var userVo systemRes.User
	err = global.LMS_DB.Where("username=?", user.Username).First(&userSelect).Error
	var decode string
	// 判断是否需要http传输的是否是加密密码，是则需要解密出密码明文
	if global.LMS_CONFIG.System.TransferPwDecode {
		// 对传输加密的密码解密，解出明文
		decode, err = utils.PasswordDecode(user.LoginTime, user.Password)
		if err != nil {
			global.LMS_LOG.Errorf("发生错误! Err: %s", err.Error())
			return userRes, err
		}
	} else {
		decode = user.Password
	}
	// 与数据库密码进行加密比对
	if !utils.IsBlankString(user.Password) && utils.BcryptCheck(decode, userSelect.Password) {
		userJson, _ := json.Marshal(userSelect)
		// 获取token  jwt
		token, err := utils.MakeToken(userSelect)
		if err != nil {
			global.LMS_LOG.Errorf("发生错误! Err: %s", err.Error())
			return userRes, err
		}
		// 获取有效时间，将token存入redis中
		ep, err := utils.ParseTimeDur(global.LMS_CONFIG.JWT.ExpiresTime)
		if err != nil {
			global.LMS_LOG.Errorf("发生错误! Err: %s", err.Error())
			return userRes, err
		}
		err = global.LMS_REDIS.Set(context.TODO(), token, userJson, ep).Err()
		if err != nil {
			global.LMS_LOG.Errorf("发生错误! Err: %s", err.Error())
			return userRes, err
		}
		userVo.RoleEn = userSelect.RoleEn
		userVo.Id = userSelect.ID
		userVo.Username = userSelect.Username
		userVo.Authorization = token
		return userVo, nil
	}
	return userRes, errors.New("验证不通过。")
}

func (u *UserService) Logout(token string) error {
	res, err := global.LMS_REDIS.Del(context.TODO(), token).Result()
	if err != nil {
		global.LMS_LOG.Errorf("发生错误! Err: %s", err.Error())
		return err
	}
	if res != 1 {
		return errors.New("退出登录失败！")
	}
	return nil
}

func (u *UserService) Register(user model.User) error {
	user.ID = utils.GetRandomId()
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	// 判断是否需要http传输的是否是加密密码，是则需要解密出密码明文
	if global.LMS_CONFIG.System.TransferPwDecode {
		// 对传输加密的密码解密，解出明文
		decode, err := utils.PasswordDecode(user.LoginTime, user.Password)
		if err != nil {
			global.LMS_LOG.Errorf("发生错误! Err: %s", err.Error())
			return err
		}
		user.Password = decode
	}
	user.Password = utils.BcryptHash(user.Password)
	err := global.LMS_DB.Create(&user).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误! Err: %s", err.Error())
		return err
	}
	return nil
}

func (u *UserService) GetUserInfo(id string) (systemRes.UserInfo, error) {
	var userInfo systemRes.UserInfo
	err := global.LMS_DB.Model(model.User{}).Where("id = ?", id).Find(&userInfo).Error
	if err != nil {
		return systemRes.UserInfo{}, err
	}
	return userInfo, nil
}

func (u *UserService) UpdateUserInfo(userInfo request.UserInfo) error {
	j, _ := json.Marshal(userInfo)
	var user model.User
	_ = json.Unmarshal(j, &user)
	user.UpdatedAt = time.Now()
	err := global.LMS_DB.Model(&model.User{}).Where("id = ?", userInfo.Id).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ChangeUserPassword(id string, pw string, changeTime time.Time) error {
	// 判断是否需要http传输的是否是加密密码，是则需要解密出密码明文
	if global.LMS_CONFIG.System.TransferPwDecode {
		// 对传输加密的密码解密，解出明文
		decode, err := utils.PasswordDecode(changeTime, pw)
		if err != nil {
			global.LMS_LOG.Errorf("发生错误! Err: %s", err.Error())
			return err
		}
		pw = decode
	}
	bcryptPw := utils.BcryptHash(pw)
	cupMap := make(map[string]interface{})
	cupMap["password"] = bcryptPw
	cupMap["updatedAt"] = time.Now()
	err := global.LMS_DB.Model(model.User{}).Where("id = ?", id).Updates(cupMap).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) UpdateUserRole(roleUr request.UserRole) error {
	err := global.LMS_DB.Model(model.User{}).Where("id = ?", roleUr.Id).Updates(roleUr).Error
	if err != nil {
		return err
	}
	return nil
}
