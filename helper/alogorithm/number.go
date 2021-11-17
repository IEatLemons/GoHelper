package algorithm

import (
	"math"
	"strconv"
)

const (
	Positive = "Positive Order"
	Reverse  = "Reverse order"
)

func BubbleSort(arr []int, Order string) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			switch Order {
			case Positive:
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			case Reverse:
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	return arr
}

func StrToFloat64Round(str string, prec int, round bool) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return Precision(f, prec, round)
}

func Precision(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		return math.Trunc(f+0.5/pow10_n) * pow10_n / pow10_n
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}
