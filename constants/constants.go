package constants

import (
	"fmt"
	"math"
)

const s string = "constant"

//PrintConstants if for printing constants
func PrintConstants() {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
