package binary

import (
	"fmt"
	"math"
	"sort"
)

func BinarySearchSimple(a []int, n int) bool {
	first := 0
	last := len(a) - 1
	for last >= first {
		mid := math.Floor(float64(first) + float64(last))
		if a[int(mid)] == n {
			return true
		} else if n < a[int(mid)] {
			last = int(mid) - 1
		} else {
			last = int(mid) + 1
		}
	}
	return false
}

func BinarySearchWithSort(list []string, word string) bool {
	i := sort.Search(len(list), func(i int) bool { return list[i] >= word })
	if i < len(list) && list[i] == word {
		fmt.Printf("The value %v was found at index: %v \n", word, i)
		return true
	} else {
		fmt.Printf("The value %v was not found in array...", word)
		return false
	}
}
