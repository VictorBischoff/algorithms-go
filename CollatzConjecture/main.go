package main

import (
	"fmt"
	//"math"
)

func collatz(startingPoint int64) []int64 {
	data := []int64{}
	for {
		data = append(data, startingPoint)
		if startingPoint%2 != 0 {
			startingPoint *= 3
			startingPoint++
		} else if startingPoint%2 == 0 {
			startingPoint /= 2
		}
		if startingPoint == 1 {
			data = append(data, startingPoint)
			break
		}
	}

	return data
}

func main() {
	//inputNum := math.Pow(10.0, 100.0)
	numberarray := []int64{}
	numberarray = append(numberarray, collatz(int64(27))...)
	fmt.Println(numberarray)
}
