package goGORMAssociation

import "gorm.io/gorm"

type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserId uint
}

func HasMany(db *gorm.DB) {
	db.AutoMigrate(&User{}, &CreditCard{})
}
