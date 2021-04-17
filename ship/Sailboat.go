package ship

type Sailboat struct {
	name         string
	shipType     string
	engineStatus int16
	speed        float64
	latitude     float64
	longitude    float64
}

func (s Sailboat) GetName() string {
	return s.name
}

func (s *Sailboat) SetName(name string) {
	s.name = name
}

func (s Sailboat) Type() string {
	return s.shipType
}

func (s Sailboat) EngineStatus() int16 {
	return s.engineStatus
}

func (s Sailboat) Speed() float64 {
	return s.speed
}

func (s Sailboat) Latitude() float64 {
	return s.latitude
}

func (s Sailboat) Longitude() float64 {
	return s.longitude
}

func NewSailboat(name string) *Sailboat {
	sailboat := &Sailboat{name: name, shipType: "Sailboat", engineStatus: OFF, speed: 0}
	return sailboat
}
