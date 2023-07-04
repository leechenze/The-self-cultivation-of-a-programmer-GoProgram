package goGORMAssociation

import (
	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	Email   string
	User1Id int
}

// User1 拥有并属于多种 Language，使用 `user_languages`连接表
type User1 struct {
	gorm.Model
	Name            string
	BillingAddress  Address
	ShippingAddress Address
	Emails          []Email
	Languages       []Language2 `gorm:"many2many:user_languages"`
}

type Language2 struct {
	gorm.Model
	Name   string
	Users1 []*User1 `gorm:"many2many:user_languages"`
}

type Address struct {
	gorm.Model
	Address string
	User1Id int
}

func EntityAssociation(db *gorm.DB) {

	// 创建这几张表
	// db.AutoMigrate(&User1{}, &Email{}, &Address{}, &Language2{})

	// 给User表创建数据
	// user := User1{
	// 	Name:            "leechenze",
	// 	BillingAddress:  Address{Address: "Billings Address - Address 1"},
	// 	ShippingAddress: Address{Address: "Shipping Address - Address 1"},
	// 	Emails: []Email{
	// 		{Email: "leeczyc@gmail.com"},
	// 		{Email: "leeczyc@163.com"},
	// 	},
	// 	Languages: []Language2{
	// 		{Name: "English"},
	// 		{Name: "Chinese"},
	// 	},
	// }
	// db.Create(&user)
	// 如果是更新一条数据记录，应该使用 FullSaveAssociations
	// db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	// 查找关联
	// var user1 User1
	// var languages []Language2
	// db.First(&user1)
	// db.Model(&user1).Association("Languages").Find(&languages)
	// for _, language := range languages {
	// 	fmt.Printf("language: %v\n", language)
	// }

	// 添加关联
	// 替换关联
	// 删除关联
	// 清空关联
	// 关联计数
	// 以上操作不在演示，请自行查阅文档。

}
