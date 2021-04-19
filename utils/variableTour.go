package utils

import (
	"fmt"
	"strconv"
)

func Swap(x, y string) (string, string) {
	return y, x
}

//println is fine for basic operations, but fmt.Println is the right way to do it
func StringsTest() {
	var date int = 16
	var month int = 12
	year := 99
	converted := strconv.Itoa(year) //strings are bytes arrays, so each byte is treated like that, a NUMBER in the ascci code
	// You can't modify a position in the string, but you can access to it
	fmt.Printf("%v/%v/%v \n", date, month, string(converted[0])) //byte converted to letter
	fmt.Printf("%v/%v/%v \n", date, month, converted[0])         // byte as ascii code
	MathOperations()
	//booleanTest()
	//bitShifting()
}
func MathOperations() {

	var first float32 = 16.5
	var second float32 = 12.4

	fmt.Printf("%.2f \n", (first + second))
	fmt.Printf("%.3f \n", (float64(first) / float64(second)))
	// Error: print((floa64(first) / second))
	// println((float64(first) / float64(second)) -> +1.330645e+000 == Not good
	// I can create a complex number (real + imaginary -> 5+8i) buuuut I don't wanna do it now
}
func BooleanTest() {
	println("---- Booleans --")
	b1 := 8                       // 1000
	b2 := 5                       // 0101
	b3 := 3                       // 0011
	println(b1 & b2)              // add a bit in a position when both bits are on 0000
	println(b1 | b2)              // add a bit in a position where a bit is on (or just an adding operation) 1101
	println(b1^b2, "--", b2^b3)   // add a bit in a position when a bit is "on" JUST in one byte. 1101 (for b1 ^b2), 0110 (for b2^b3)
	println(b1&^b2, "--", b2&^b3) // add a bit in a position when a bit in the first byte is ON and the same position in the second byte is OFF. 1000 (for b1 ^b2), 0100 (for b2^b3)

}

func BitShifting() {
	println("---- bitShifting ---")
	a := 16         // 2^4
	println(a >> 2) // 2^4 / 2^2 -> 16/4
	println(a << 3) // 2^4 * 2^3 -> 16*8
}
