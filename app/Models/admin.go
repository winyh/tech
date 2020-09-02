package Models

import (
	. "SJService/database"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}

// Store 新增admin用户
func (admin *Admin) Store() (userID uint, err error) {
	result := DB.Create(&admin)
	userID = admin.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

// Destroy 删除admin用户
func (admin *Admin) Destroy(id int) (err error) {
	// result := DB.Delete(&admin, id) // 软删除
	result := DB.Unscoped().Delete(&admin, id)  // 永久删除
	if result.Error != nil {
		err = result.Error
	}
	return
}


// Update 修改admin用户
func (admin *Admin) Update(id int) (err error) {
	result := DB.Model(&admin).Where("id = ?", id).Updates(&admin)
	if result.Error != nil {
		err = result.Error
	}
	return
}

// FindOne 查询admin用户
func (admin *Admin) FindOne(id int) (user Admin, err error) {
	result := DB.First(&admin, id)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// FindAll 查询admin用户
func (admin *Admin) FindAll() (Admin []Admin, err error) {

	result := DB.Find(&Admin) // 这里的 &Admin 跟返回参数要一致

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}