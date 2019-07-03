package recursion

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//PrintRecursion is for recursion example.
func PrintRecursion() {
	fmt.Println(fact(7))
}
