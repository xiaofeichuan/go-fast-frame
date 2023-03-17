package service

import (
	"errors"
	"go-fast-frame/app/dto"
	"go-fast-frame/common/utils"
	"go-fast-frame/global"
	"go-fast-frame/model"
)

type SysUserService struct{}

// Query 用户分页查询
func (s *SysUserService) Query(query dto.SysUserQuery) (list []dto.SysUserVo, total int64, err error) {
	db := global.DB.Model(&model.SysUser{})
	//查询条件
	if query.UserName != "" {
		value := "%" + query.UserName + "%"
		db = db.Where("nick_name LIKE ? or user_name like ? or phone like ?", value, value, value)
	}

	//总条数
	db.Count(&total)

	offset := (query.PageNum - 1) * query.PageSize
	//查询数据
	err = db.Order("id DESC").Offset(offset).Limit(query.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, err
}

// Add 添加用户
func (s *SysUserService) Add(addDto dto.SysUserAddDto) error {

	if addDto.UserName == "" {
		return errors.New("账号名称不允许为空")
	}

	var count int64
	global.DB.Model(&model.SysUser{}).Where("user_name = ?", addDto.UserName).Count(&count)
	if count != 0 {
		return errors.New("账号名称已存在")
	}

	//密码盐
	var salt = utils.Text.GetRoundNumber(15)

	//加密 => MD5（密码+密码盐）
	var pwd = utils.Encryption.Md5(addDto.Password + salt)

	var user = &model.SysUser{
		UserName: addDto.UserName,
		NickName: addDto.NickName,
		UserType: 0,
		Email:    addDto.Email,
		Phone:    addDto.Phone,
		Password: pwd,
		Salt:     salt,
		Gender:   addDto.Gender,
		Status:   addDto.Status,
		Remark:   addDto.Remark,
	}
	err := global.DB.Create(user).Error
	return err
}

// Update 更新用户
func (s *SysUserService) Update(updateDto dto.SysUserUpdateDto) error {
	err := global.DB.Model(&model.SysUser{}).Where("id = ?", updateDto.Id).Updates(map[string]interface{}{
		"nick_name": updateDto.NickName,
		"user_type": 0,
		"email":     updateDto.Email,
		"phone":     updateDto.Phone,
		"gender":    updateDto.Gender,
		"status":    updateDto.Status,
		"remark":    updateDto.Remark,
	}).Error
	return err
}

// Delete 删除用户
func (s *SysUserService) Delete(id int64) error {
	err := global.DB.Delete(&model.SysUser{}, "id = ?", id).Error
	return err
}

// Detail 获取用户详情
func (s *SysUserService) Detail(id int64) (vo dto.SysUserVo, err error) {
	err = global.DB.Model(&model.SysUser{}).Where("id = ?", id).Scan(&vo).Error
	return vo, err
}

// List 用户列表
func (s *SysUserService) List() (vos []dto.SysUserVo, err error) {
	err = global.DB.Model(&model.SysUser{}).Scan(&vos).Error
	return vos, err
}
