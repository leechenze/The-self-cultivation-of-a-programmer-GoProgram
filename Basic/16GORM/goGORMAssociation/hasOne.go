package goGORMAssociation

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Number string
	UserId int
}

type Owner struct {
	gorm.Model
	CreditCard CreditCard
}

func HasOne(db *gorm.DB) {
	// TODO
}
