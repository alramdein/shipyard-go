package ship

type Yacht struct {
	name         string
	shipType     string
	engineStatus int16
	speed        float64
	latitude     float64
	longitude    float64
}

func (y Yacht) GetName() string {
	return y.name
}

func (y *Yacht) SetName(name string) {
	y.name = name
}

func (y Yacht) Type() string {
	return y.shipType
}

func (y Yacht) EngineStatus() int16 {
	return y.engineStatus
}

func (y Yacht) Speed() float64 {
	return y.speed
}

func (y Yacht) Latitude() float64 {
	return y.latitude
}

func (y Yacht) Longitude() float64 {
	return y.longitude
}

func NewYacht(name string) *Yacht {
	yacht := &Yacht{name: name, shipType: "Yacht", engineStatus: OFF, speed: 0}
	return yacht
}
