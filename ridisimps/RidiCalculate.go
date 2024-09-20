package ridisimps

func SumIntSlices(s [][]int) int {
	cost := 0
	for _, slice := range s {
		cost += SumIntSlice(slice)
	}

	return cost
}

func SumIntSlice(s []int) int {
	total := 0

	for _, value := range s {
		total += value
	}

	return total
}

func PutCommasAtNumber(number string) string {
	byteNumber := []byte(number)
	length := len(number)
	mod := length % 3

	if length <= 3 {
		return number
	}

	result := make([]byte, 0, 9)
	var start int
	if mod > 0 {
		result = append(result, byteNumber[:mod]...)
		start = mod
	} else {
		result = append(result, byteNumber[:3]...)
		start = 3
	}

	for i := start; i <= length-3; i += 3 {
		result = append(result, ',')
		result = append(result, byteNumber[i:i+3]...)
	}

	return string(result)
}

func IsSliceEmpty(s []int) bool {
	return len(s) == 0
}
