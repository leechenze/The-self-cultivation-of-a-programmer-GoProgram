package goGORMActionFunc

import (
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) {
	// var user User
	// var users []User

	// 删除一条记录
	// db.Find(&user, "name = ?", "tom")
	// db.Delete(&user)

	// 删除一条记录（另一种方式）
	// db.Where("name = ?", "Michele").Delete(&User{})

	// 根据主键ID删除
	// db.Delete(&user, "name = ?", "Elizabeth")

	// 物理删除，前面删除都是软删除。
	// db.Unscoped().Where("name", "tom").Delete(&User{})

}
