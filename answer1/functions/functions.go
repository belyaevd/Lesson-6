package functions

func sum(arr []float64) float64 {
	total := float64(0)
	for _, x := range arr {
		total += x
	}
	return total
}
