package goGORMAssociation

import "gorm.io/gorm"

// 信用卡和用户
// type CreditCard struct {
// 	gorm.Model
// 	Number  string
// 	OwnerId int
// }
//
// type Owner struct {
// 	gorm.Model
// 	CreditCard CreditCard
// }

// 主人和猫狗宠物（多态关联）
type Cat struct {
	ID   int
	Name string
	Pet  Pet `gorm:"polymorphic:Master;"`
}

type Dog struct {
	ID   int
	Name string
	Pet  Pet `gorm:"polymorphic:Master;"`
}

type Pet struct {
	ID         int
	Name       string
	MasterID   int
	MasterType string
}

func HasOne(db *gorm.DB) {
	// 创建 Owner 和 CreditCarad 的表(注意 AutoMigrate 的先后顺序)
	// db.AutoMigrate(&Owner{}, &CreditCard{})

	// 创建 cat 和 dog 和 pet 的表
	// db.AutoMigrate(&Pet{}, &Dog{}, &Cat{})

	// 创建一条Dog表的数据
	db.Create(&Dog{Name: "ethan", Pet: Pet{Name: "pet1"}})
	// 创建一条Cat表的数据
	db.Create(&Cat{Name: "robin", Pet: Pet{Name: "pet2"}})

}
