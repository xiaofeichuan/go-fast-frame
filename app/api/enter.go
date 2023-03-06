package api

import (
	"go-fast-frame/app/service"
)

var (
	authService = &service.SysAuthService{}
	userService = &service.SysUserService{}
)
