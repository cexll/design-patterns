package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyWeight(t *testing.T) {
	game := newGame()

	// Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	// Add CounterTerrorist
	game.addCounterTerrorists(CounterTerroristDressType)
	game.addCounterTerrorists(CounterTerroristDressType)
	game.addCounterTerrorists(CounterTerroristDressType)

	dressFactorySingleInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactorySingleInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor:%s\n", dressType, dress.getColor())
	}
}
