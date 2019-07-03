package multiplereturns

import "fmt"

func vals() (int, int) {
	return 3, 7
}

//PrintMultiple is for multiple return example
func PrintMultiple() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
