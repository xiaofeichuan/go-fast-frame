package api

import (
	"go-fast-frame/app/dto"
	"go-fast-frame/common/dto/request"
	"go-fast-frame/common/dto/response"

	"github.com/gin-gonic/gin"
)

type SysUserApi struct{}

// Query
// @Tags 用户
// @Summary 用户查询
// @Security ApiKeyAuth
// @Param data query dto.SysUserQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysUserVo}}
// @Router /system/user/query [get]
func (s *SysUserApi) Query(c *gin.Context) {
	var query dto.SysUserQuery
	c.ShouldBindQuery(&query)
	list, total, err := userService.Query(query)
	response.Complete(response.PageResult{List: list, TotalCount: total}, err, c)
}

// Add
// @Tags 用户
// @Summary 添加用户
// @Security ApiKeyAuth
// @Param data body dto.SysUserAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/user/add [post]
func (s *SysUserApi) Add(c *gin.Context) {
	var addDto dto.SysUserAddDto
	c.ShouldBindJSON(&addDto)
	err := userService.Add(addDto)
	response.Complete(nil, err, c)
}

// Update
// @Tags 用户
// @Summary 更新用户
// @Security ApiKeyAuth
// @Param data body dto.SysUserUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/user/update [post]
func (s *SysUserApi) Update(c *gin.Context) {
	var updateDto dto.SysUserUpdateDto
	c.ShouldBindJSON(&updateDto)
	err := userService.Update(updateDto)
	response.Complete(nil, err, c)
}

// Delete
// @Tags 用户
// @Summary 删除用户
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/user/delete [post]
func (s *SysUserApi) Delete(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)
	err := userService.Delete(idInfoDto.Id)
	response.Complete(nil, err, c)
}

// Detail
// @Tags 用户
// @Summary 获取用户详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "用户id"
// @Success 200 {object} response.JsonResult{data=dto.SysUserVo}
// @Router /system/user/detail [get]
func (s *SysUserApi) Detail(c *gin.Context) {
	var idInfoDto request.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	obj, err := userService.Detail(idInfoDto.Id)
	response.Complete(obj, err, c)
}

// List
// @Tags 用户
// @Summary 用户列表
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.SysUserVo}
// @Router /system/user/list [get]
func (s *SysUserApi) List(c *gin.Context) {
	objs, err := userService.List()
	response.Complete(objs, err, c)
}
