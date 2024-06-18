package foohandler

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func (e gasEngine) milesLeft() float32 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() float32 {
	return e.kwh * e.mpkwh
}

func (e car[gasEngine]) milesLeft() float32 {
	return e.engine.milesLeft() // ???
}

func main() {
	// Summing a slice...
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	// A car
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}

	gasCar.engine.milesLeft()

	fmt.Printf("Miles left on gas: %v\n", gasCar.milesLeft())

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine:   electricEngine{57.5, 4.17},
	}

	fmt.Printf("Miles left on electric car: %v\n", electricCar.milesLeft())

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
