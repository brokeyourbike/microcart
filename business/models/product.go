package models

type Product struct {
	ID     uint64 `gorm:"primaryKey"`
	ShopID uint64 `gorm:"index"`
}
