package goGORMAssociation

import "gorm.io/gorm"

// user_languages 是多对多关系的连接表或者说是中间表
type People struct {
	gorm.Model
	UserName  string
	Languages []Language `gorm:"many2many:people_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

// 反向引用
type People1 struct {
	gorm.Model
	UserName   string
	Languages1 []*Language1 `gorm:"many2many:people_languages;"`
}

type Language1 struct {
	gorm.Model
	Name    string
	People1 []*People1 `gorm:"many2many:people_languages;"`
}

func ManyToMany(db *gorm.DB) {
	// 创建People和Language表
	// db.AutoMigrate(&People{}, &Language{})

	// 创建一条People表和Language表的记录
	people := People{UserName: "Nixon"}
	language := Language{Name: "English"}
	db.Create(&people)
	db.Create(&language)

}
