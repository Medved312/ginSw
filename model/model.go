package model

type Product struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Categories  []*Category `gorm:"many2many:productCategory"`
}
type Category struct {
	Id      uint `gorm:"primaryKey"`
	Name    string
	Product []*Product `gorm:"many2many:productCategory"`
}

type Order struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	quantity uint
}
