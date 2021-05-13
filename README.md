# Shipyard-go
Ship package can create ship instances. Currently it has 3 instance: Motorboat, Sailboat and Yacht. You can update, customize or add it by yourself by clone the project. 

## Requirements
Go installed on your machine. More [install Go.](https://golang.org/doc/install)

## Installation
```go
go get -v github.com/alramdein/shipyard-go/ship
```

## Usage 
Import the `Ship` package
```golang
import (
	"github.com/alramdein/shipyard-go/ship"
)
```

Implement
```golang
type Shiper ship.Shiper 

func main() {
    motorboat = ship.NewMotorboat("MB-01")
    sailboat = ship.NewSailboat("SB-01")
    yacht = ship.NewMotorboat("YC-01")

    fmt.Println("Motorboat-1 name: " + motorboat1.GetName())
    fmt.Println()

    motorboat1.SetName("MB-01.11")
    fmt.Println("Motorboat-1 name after SetName: " + motorboat1.GetName())
}

```
##  Available functions
### Public interface function s
1. `NewMotorboat(name)` to create new motorboat instance. 
2. `NewSailboat(name)` to create new sailboat instance. 
3. `NewYacht(name)` to create new yacht instance. 
### Public object functions
1. `GetName()` string
2. `SetName(name string)`
3. `Type()` string
4. `EngineStatus()` int16
5. `Speed()` float64
6. `Latitude()` float64
7. `Longitude()` float64

And it is all accept a `string` parameter of its ship name.


## Modification
You can customize it by yourself by cloning it with 
```git
git clone https://github.com/alramdein/shipyard-go.git
```

### Run
On root project, run `go run main.go`

**OR**

On root project, run `go build`
And then run the executable file with `shipyard-go.exe`

> Note: Sometimes it took quite long at first `run` because it needs to download the package for the repository

## Author
This project original created from scratch by Alif Ramdani (alramdein). Please put me on credit when you use it publicly or commercially.