package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	myPic := make([][]uint8, dy)
	for x := 0; x < dx; x++ {
		myPic[x] = make([]uint8, dx)
		for y := 0; y < dy; y++ {
			data := (x + y) / 2
			myPic[x][y] = uint8(data)
		}
	}
	return myPic //This do not work on windows
}

func WordCount(s string) map[string]int {

	return map[string]int{"x": 1}

}

func main() {
	fmt.Println(("hola como estás soy la loca"))
	wc.Test
}
