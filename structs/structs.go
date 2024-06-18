package main

import (
	"fmt"
	"time"
)

type engine interface {
	milesLeft() uint8
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint8) {
	if miles < e.milesLeft() {
		fmt.Println("You can make it.")
	} else {
		fmt.Println("Need to fuel up first it.")
	}
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println("In this tutorial we are going to look into structs and interfaces.")

	myEngine := electricEngine{25, 15, owner{"Lars"}}
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	canMakeIt(myEngine, 110)

	// Channels
	var c = make(chan int, 5)
	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}
