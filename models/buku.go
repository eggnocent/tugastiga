package models

type Buku struct {
	ID        uint   `gorm:"primaryKey"`
	NamaBuku  string `gorm:"not null"`
	Penulis   string `gorm:"not null"`
	TglTerbit string `gorm:"not null"`
	UserID    uint   `gorm:"unique"`
	User      *User  `gorm:"foreignKey:UserID"`
}
