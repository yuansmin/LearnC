package sort

func Sort(data []int) []int {
	length := len(data)
	for j := length; j > 0; j-- {
		// the array is sorted if no swaped happened in one bubble loop
		swaped := false
		for i := 1; i < j; i++ {
			if data[i-1] > data[i] {
				data[i-1], data[i] = data[i], data[i-1]
				swaped = true
			}
		}
		if !swaped {
			return data
		}
	}
	return data
}
