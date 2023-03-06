package api

import (
	"go-fast-frame/app/dto"
	"go-fast-frame/common/dto/response"

	"github.com/gin-gonic/gin"
)

type SysAuthApi struct{}

// Login
// @Tags 授权
// @Summary 用户登录
// @Security ApiKeyAuth
// @Param data body dto.LoginDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/auth/login [post]
func (s *SysAuthApi) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	c.ShouldBindJSON(&loginDto)
	token, err := authService.Login(loginDto)
	response.Complete(token, err, c)
}
