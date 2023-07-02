package goGORMActionFunc

import (
	"gorm.io/gorm"
	"time"
)

// 声明模型
type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time `gorm:"default:nil"`
}

// 钩子函数
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	println("beforeCreate is Running")
	return
}

func Create(db *gorm.DB) {
	// 创建User表（表明就是模型名称加s后缀），创建一次即可，这里先注释了。
	// db.AutoMigrate(&User{})

	// user := User{
	// 	Name:     "Nixon",
	// 	Age:      20,
	// 	Birthday: time.Now(),
	// }

	// 创建记录（插入一条数据）
	// tx := db.Create(&user)
	// fmt.Printf("tx.RowsAffected: %v\n", tx.RowsAffected)
	// fmt.Printf("user: %v\n", user)

	// 用指定的字段创建记录。（只添加指定的Name和Age字段）
	// db.Select("Name", "Age").Create(&user)

	// 批量插入
	// var users = []User{
	// 	{Name: "Hillary"},
	// 	{Name: "Clinton"},
	// 	{Name: "Obama"},
	// 	{Name: "Roosevelt"},
	// }
	// db.Select("Name", "Age").Create(&users)

	// 根据Map创建
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "BlinKen", "Age": 40,
	})

}
