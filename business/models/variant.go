package models

type Variant struct {
	ID        uint64 `gorm:"primaryKey"`
	ProductID uint64 `gorm:"index"`
}
