/* Program: Shipyard-go */
/* Author: Alif Ramdani */
/* Social Media: alramdein */
/* Date: 17/04/2021 */

package main

import (
	"fmt"

	"github.com/alramdein/shipyard-go/ship"
)

const MOTORBOAT = 1
const SAILBOAT = 2
const YACHT = 3

type Shiper ship.Shiper

func NewShip(shipType int, name string) Shiper {
	var newShip Shiper
	switch shipType {
	case MOTORBOAT:
		{
			newShip = ship.NewMotorboat(name)
		}
	case SAILBOAT:
		{
			newShip = ship.NewSailboat(name)
		}
	case YACHT:
		{
			newShip = ship.NewYacht(name)
		}
	}
	return newShip
}

func main() {
	var motorboat1 Shiper = NewShip(MOTORBOAT, "MB-01")
	var motorboat2 Shiper = NewShip(MOTORBOAT, "MB-02")
	fmt.Println("Motorboat-1 name: " + motorboat1.GetName())
	fmt.Println("Motorboat-2 name: " + motorboat2.GetName())
	fmt.Println()

	motorboat1.SetName("MB-01.11")
	fmt.Println("Motorboat-1 name after SetName: " + motorboat1.GetName())
	fmt.Println()

	var sailboat Shiper = NewShip(SAILBOAT, "SB-01")
	var yacht Shiper = NewShip(YACHT, "YC-02")
	fmt.Printf("Sailboat-1 engine status: %d \n", sailboat.EngineStatus())
	fmt.Printf("Yacht-1 speed: %f Km/h \n", yacht.Speed())
}
