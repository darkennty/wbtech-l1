package main

import (
	"fmt"
	"math"
)

func main() {
	tArray := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tMap := make(map[int][]float64)

	for _, t := range tArray {
		var key int
		if t >= 0 {
			key = int(floorToTen(t))
		} else {
			key = int(ceilToTen(t))
		}

		tMap[key] = append(tMap[key], t)
	}

	for k, v := range tMap {
		fmt.Printf("%d %.1f\n", k, v)
	}
}

func floorToTen(num float64) float64 {
	return math.Floor(num/10) * 10
}

func ceilToTen(num float64) float64 {
	return math.Ceil(num/10) * 10
}
