package goGORMActionFunc

import (
	"gorm.io/gorm"
)

func Update(db *gorm.DB) {
	// var user User

	// 更新记录
	// db.Find(&user, 7)
	// user.Name = "Abraham Lincoln"
	// user.Age = 88
	// user.Birthday = time.Now()
	// db.Save(&user)

	// 更新记录的单个字段（&User{}是没有查出来的一个模型）
	// db.Model(&User{}).Where("name = ?", "Roosevelt").Update("Age", 66)

	// 更新多个字段（根据结构体）
	// db.Find(&user, 14)
	// db.Model(&user).Updates(User{Name: "Michele", Age: 40, Active: false})

	// 更新多个字段（根据Map）
	// db.Find(&user, 16)
	// db.Model(&user).Updates(map[string]interface{}{"name": "Sergey", "age": 23, "Active": true})

	// 批量更新

}
