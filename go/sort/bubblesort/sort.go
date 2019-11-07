package sort

func Sort(data []int) []int {
	length := len(data)
	swaped := true
	for j := length; j > 0; j-- {
		// the array is sorted if no swaped happened in pre bubble loop
		if !swaped {
			return data
		}
		swaped = false
		for i := 1; i < j; i++ {
			if data[i-1] > data[i] {
				data[i-1], data[i] = data[i], data[i-1]
				swaped = true
			}
		}
	}
	return data
}
