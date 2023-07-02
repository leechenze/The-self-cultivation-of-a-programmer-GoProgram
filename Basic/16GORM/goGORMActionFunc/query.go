package goGORMActionFunc

import (
	"fmt"
	"gorm.io/gorm"
)

func Query(db *gorm.DB) {
	// user := User{
	// 	Name:     "Roosevelt",
	// 	Age:      20,
	// 	Birthday: time.Now(),
	// }

	// 获取第一条记录，按主键ID升序
	// db.First(&user)
	// fmt.Printf("user: %v\n", user)

	// 获取第一条记录，不指定排序字段
	// db.Take(&user)
	// fmt.Printf("user: %v\n", user)

	// 获取最后一条记录，按主键ID降序
	// db.Last(&user)
	// fmt.Printf("user: %v\n", user)

	// 根据ID查询
	// db.Find(&user, 13)
	// fmt.Printf("user: %v\n", user)

	// 根据多个ID查询
	// var users []User
	// db.Find(&users, []int{12, 13, 14})
	// for _, user := range users {
	// 	fmt.Printf("user: %v\n", user)
	// }

	// 查找所有记录
	// var users []User
	// db.Find(&users)
	// fmt.Printf("users: %v\n", users)

	// 根据条件查找记录
	// var user User
	var users []User
	// db.Where("name = ?", "Blinken").Find(&user)
	// 以结构体为条件查询
	// fmt.Printf("user: %v\n", user)
	// db.Where(&User{Name: "Nixon"}).Find(&user)
	// 以Map为条件查询。
	// fmt.Printf("user: %v\n", user)
	// db.Where(map[string]interface{}{"Name": "Nixon"}).Find(&user)
	// 以ID切片为条件的查询（多个ID查询）
	db.Where([]int{12, 13, 14}).Find(&users)
	fmt.Printf("users: %v\n", users)

	// 等等还有很多 Or, Not, Order(排序), Limit, Offset(这来一般分页用), Distinct(从多条记录中取一条), Join
	// Scan 和 Find类似，只不过Scan将结果放到了结构体中了，find是放在了结构体的实例对象中了。

}
