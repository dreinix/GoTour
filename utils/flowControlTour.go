package utils

import (
	"fmt"
	"math"
	"runtime"
)

func Sqrt(x float64) float64 {
	var z float64 = x / 3
	var tmp float64 = z * 2
	for math.Abs(tmp-z) > 1e-8 {
		tmp = z
		z -= (z*z - x) / (2 * z)
	}
	return z

}

func Swtch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
