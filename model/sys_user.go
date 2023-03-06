package model

import (
	"go-fast-frame/constant"

	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UserName string              //用户账号
	NickName string              //用户昵称
	UserType constant.UserType   //用户类型（0普通账号，1超级管理员）
	Email    string              //用户邮箱
	Phone    *string             //手机号码
	Password string              //密码
	Salt     string              //密码盐
	Sex      int                 //用户性别（0未知，1男，2女）
	Avatar   *string             //头像地址
	Status   constant.UserStatus //帐号状态（0正常 1停用）
	Remark   *string             //备注
}
