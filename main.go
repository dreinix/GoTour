package main

import (
	"github.com/dreinix/GoTour/utils"
)

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

func main() {
	utils.Show(Pic)
}
