package linear

import (
	"fmt"
	"strings"
)

func SimpleLinearSearch(list []int, n int) bool {
	for index, i := range list {
		if i == n {
			fmt.Printf("found value %v in the list at index %v \n", n, index)
			return true
		}
	}
	return false
}

func FindLetterInString(s string, c string) bool {
	arr := strings.Split(s, "")
	for index, i := range arr {
		if i == c {
			fmt.Printf("Found: %v, at index %v \n", c, index)
			return true
		}
	}
	return false
}
