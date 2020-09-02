package Models

import (
	. "SJService/database"
	"gorm.io/gorm"
)

type Admins struct {
	gorm.Model
	Name string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}

// Store 新增admin用户
func (admin *Admins) Store() (userID uint, err error) {
	result := DB.Create(&admin)
	userID = admin.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

// Destroy 删除admin用户
func (admin *Admins) Destroy() (err error) {

	result := DB.Delete(&admin)
	if result.Error != nil {
		err = result.Error
	}
	return
}


// Update 修改admin用户
func (admin *Admins) Update(id int64) (user Admins, err error) {
	result := DB.Model(&admin).Where("id = ?", id).Updates(&admin)
	if result.Error != nil {
		err = result.Error
	}
	return
}

// FindOne 查询admin用户
func (admin *Admins) FindOne(id int) (user Admins, err error) {
	result := DB.First(&admin, id)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// FindAll 查询admin用户
func (admin *Admins) FindAll() (admins []Admins, err error) {

	result := DB.Find(&admins) // 这里的 &admins 跟返回参数要一致

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}