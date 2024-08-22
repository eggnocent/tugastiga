package models

type Penulis struct {
	ID           uint   `gorm:"primaryKey"`
	NamaPenulis  string `gorm:"not null"`
	EmailPenulis string `gorm:"not null;unique"`
	Buku         []Buku `gorm:"foreignKey:IdPenulis"`
}
