package dto

import (
	"go-fast-frame/common/dto/request"
	"go-fast-frame/constant"
	"time"
)

// SysUserAddDto 创建
type SysUserAddDto struct {
	UserName string              `json:"userName"`  //用户账号
	NickName string              `json:"nickName" ` //用户昵称
	Email    string              `json:"email"`     //用户邮箱
	Phone    *string             `json:"phone"`     //手机号码
	Password string              `json:"password"`  //密码
	Sex      int                 `json:"sex"`       //用户性别（0未知，1男，2女）
	Status   constant.UserStatus `json:"status"`    //帐号状态（0正常 1停用）
	Remark   *string             `json:"remark"`    //备注
}

// SysUserUpdateDto 更新
type SysUserUpdateDto struct {
	Id       int64               `json:"id"`        //编号
	NickName string              `json:"nickName" ` //用户昵称
	Email    string              `json:"email"`     //用户邮箱
	Phone    *string             `json:"phone"`     //手机号码
	Sex      int                 `json:"sex"`       //用户性别（0未知，1男，2女）
	Status   constant.UserStatus `json:"status"`    //帐号状态（0正常 1停用）
	Remark   *string             `json:"remark"`    //备注
}

// SysUserQuery 查询
type SysUserQuery struct {
	request.PageQuery
	UserName string `json:"userName" form:"userName"` //用户昵称
}

// SysUserVo 输出对象
type SysUserVo struct {
	Id        int64     `json:"id"`        //编号
	UserName  string    `json:"userName"`  //用户账号
	NickName  string    `json:"nickName" ` //用户昵称
	UserType  int       `json:"userType"`  //用户类型（0普通账号，1超级管理员）
	Email     *string   `json:"email" `    //用户邮箱
	Phone     *string   `json:"phone"`     //手机号码
	Sex       int       `json:"sex"`       //用户性别（0未知，1男，2女）
	Avatar    string    `json:"avatar"`    //头像地址
	Status    int       `json:"status"`    //帐号状态（0正常 1停用）
	Remark    *string   `json:"remark"`    //备注
	CreatedAt time.Time `json:"createdAt"` //创建时间
	UpdatedAt time.Time `json:"updatedAt"` //更新时间
}
