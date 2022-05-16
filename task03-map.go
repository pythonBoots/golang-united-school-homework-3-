package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	a := []int{}
	s := []string{}
	for k, _ := range input {
		a = append(a, k)
	}
	sort.Ints(a)
	for _, n := range a {
		s = append(s, input[n])
	}
	result = s
	return
}
