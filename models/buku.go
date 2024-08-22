package models

type Buku struct {
	ID        uint    `gorm:"primaryKey"`
	NamaBuku  string  `gorm:"not null"`
	TglTerbit string  `gorm:"not null"`
	IdPenulis uint    `gorm:"not null"`
	Penulis   Penulis `gorm:"foreignKey:IdPenulis"`
	UserID    uint    `gorm:"unique"`
	User      *User   `gorm:"foreignKey:UserID"`
}
