package ship

type Shiper interface {
	GetName() string
	SetName(name string)
	Type() string
	EngineStatus() int16
	Speed() float64
	Latitude() float64
	Longitude() float64
}

const ON = 1
const OFF = 0
