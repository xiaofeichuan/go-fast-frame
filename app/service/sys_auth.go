package service

import (
	"errors"
	"go-fast-frame/app/dto"
	"go-fast-frame/common/utils"
	"go-fast-frame/constants"
	"go-fast-frame/global"
	"go-fast-frame/model"
	"time"
)

type SysAuthService struct{}

// Login 用户登录
func (s *SysAuthService) Login(loginDto dto.LoginDto) (token string, err error) {

	var user model.SysUser

	err = global.DB.Where("user_name = ?", loginDto.UserName).First(&user).Error
	if err != nil {
		return "", errors.New("账号不存在")
	}

	if user.Status == constants.UserStatusDisable {
		return "", errors.New("账号已被禁用")
	}

	pwd := utils.Encryption.Md5(loginDto.Password + user.Salt)
	if pwd != user.Password {
		return "", errors.New("密码错误")
	}

	//生成token
	var claims = utils.UserAuthClaims{
		UserId:   user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		UserType: user.UserType,
	}
	token, err = utils.Jwt.GenerateToken(claims, time.Now().AddDate(0, 0, 1))
	return token, err
}
