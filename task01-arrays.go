package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
	length := len(input)
	for _, n := range input {
		sum += n
	}
	result = sum / float32(length)
	return
}
