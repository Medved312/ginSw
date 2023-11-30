package views

type CarView struct {
	Id             uint `json:"id"`
	Price          string
	Power          int
	Consumption    float32
	Overclocking   float32
	MaxSpeed       int
	ModelID        uint
	Description    string `json:"description"`
	TransmissionID uint
}
