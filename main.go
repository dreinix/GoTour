package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/dreinix/GoTour/variableTest"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(50))
	//x := 4.68
	fmt.Printf("I have the sqare of %.3g and thats pi -> %.5f \n", math.Sqrt(12), math.Pi)
	a, x := variableTest.Swap("inix", "hola")

	println(a)
	println(x)

}
