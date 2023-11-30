package model

type Car struct {
	ID             uint `gorm:"primaryKey"`
	Price          string
	Power          int
	Consumption    float32
	Overclocking   float32
	MaxSpeed       int
	Model          Model `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ModelID        uint
	Description    string
	Transmission   Transmission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransmissionID uint
}
type Mark struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Model   Model `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ModelID uint
}

type Model struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
type Transmission struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
