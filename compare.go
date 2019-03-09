package main

func max(arr []float64) float64 {
	var maximum float64 = 0
	for _, e := range arr {
		if e > maximum {
			maximum = e
		}
	}
	return maximum
}

func min(arr []float64) float64 {
	var minimum float64 = arr[0]
	for _, e := range arr {
		if e < minimum {
			minimum = e
		}
	}
	return minimum
}