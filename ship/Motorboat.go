package ship

type Motorboat struct {
	name         string
	shipType     string
	engineStatus int16
	speed        float64
	latitude     float64
	longitude    float64
}

func (m Motorboat) GetName() string {
	return m.name
}

func (m *Motorboat) SetName(name string) {
	m.name = name
}

func (m Motorboat) Type() string {
	return m.shipType
}

func (m Motorboat) EngineStatus() int16 {
	return m.engineStatus
}

func (m Motorboat) Speed() float64 {
	return m.speed
}

func (m Motorboat) Latitude() float64 {
	return m.latitude
}

func (m Motorboat) Longitude() float64 {
	return m.longitude
}

func NewMotorboat(name string) *Motorboat {
	motorboat := &Motorboat{name: name, shipType: "Motorboat", engineStatus: OFF, speed: 0}
	return motorboat
}
