package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := &VeggeMania{}

	// add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	// add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
