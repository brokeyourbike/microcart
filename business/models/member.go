package models

import "gorm.io/datatypes"

type MemberPermission string

type Member struct {
	ID          uint64 `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Email       string
	Permissions datatypes.JSONSlice[MemberPermission]
}
