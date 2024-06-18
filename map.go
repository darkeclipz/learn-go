package foohandler

import "fmt"

type Coordinate2D struct {
	Lat, Long float64
}

var m = map[string]Coordinate2D{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
	elem, exists := m["Bell Labs"]
	if exists {
		fmt.Println(elem)
	}
}
