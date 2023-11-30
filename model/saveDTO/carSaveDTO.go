package saveDTO

type CarDTO struct {
	Price          string
	Power          int
	Consumption    float32
	Overclocking   float32
	MaxSpeed       int
	ModelID        uint
	Description    string `json:"description"`
	TransmissionID uint
}
