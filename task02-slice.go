package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	length := len(input)
	N := input
	for i := 0; i <= length/2; i++ {
		N[i], N[length-1-i] = N[length-1-i], N[i]

	}
	result = N

	return
}
