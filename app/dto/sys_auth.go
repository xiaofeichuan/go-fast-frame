package dto

type LoginDto struct {
	UserName string `json:"userName"` //用户账号
	Password string `json:"password"` // 密码
}
