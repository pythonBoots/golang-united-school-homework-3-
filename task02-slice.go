package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	length := len(input)
	result = input
	for i := 0; i <= length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]

	}

	return
}
