package utils

import (
	"fmt"
	"runtime"
)

func Sqrt(x float64) float64 {
	var z float64
	if x < 10 { //10 = 2*3+1
		z = float64(3 * x)
	} else {
		z = float64(x / 3) // factor de 3
	}

	var tmp float64 = z * 2
	for tmp-z > 0.00000001 { // 6 decimales aceptables.
		tmp = z
		z -= (z*z - x) / (2 * z)
	}
	return z

}

func swtch() {
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
