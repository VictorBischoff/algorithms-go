package recursion

import (
	"fmt"
	"math"
)

// CountTo100 counts from the given number (n) up to 100
func CountTo100(n int) int {
	if n <= int(math.Pow(10.00, 2.00)) {
		fmt.Println(n)
		return CountTo100(n + 1)
	} else {
		return n
	}

}
