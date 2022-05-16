package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	length := len(input)
	for _, n := range input {
		result += n
	}
	result = result / float32(length)
	return
}
