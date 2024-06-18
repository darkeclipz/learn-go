package main
import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	var intNum int = 10
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var r, g, b uint8 = 0xff, 0x55, 0x11
	fmt.Printf("(%d, %d, %d)", r, g, b)
	fmt.Println()

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool = false
	fmt.Println(myBool)

	myVar := 10
	fmt.Println(myVar)

	fmt.Println(double(10))

	num := 10
	den := 22
	var result, remainder, err = intDivision(num, den)
	if err!=nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Division of %d\\%d has resulted in %d with remainder %d.", num, den, result, remainder);
	}
	fmt.Println()

	intArr := [...]int32{1,2,3,4,5}
	fmt.Println(intArr[1:3])
}

func double(x int) int {
	return x * x
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by Zero.")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}


