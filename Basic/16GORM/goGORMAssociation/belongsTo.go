package goGORMAssociation

import "gorm.io/gorm"

type Company struct {
	Id   int
	Name string
}

type Empolyee struct {
	gorm.Model
	Name      string
	CompanyId int
	Company   Company
}

func BelongsTo(db *gorm.DB) {
	// 创建Company 和 Empolyee的表（公司和员工表）
	db.AutoMigrate(&Company{}, &Empolyee{})
}
