package AocHelper

import "strconv"

func (in input) Raw() string {
	return in.rawInput
}

func (in input) Strings() []string {
	return getLines(in.rawInput)
}

func (in input) Ints() ([]int, error) {
	var intArr []int

	for _, v := range getLines(in.rawInput) {
		intV, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		intArr = append(intArr, intV)
	}

	return intArr, nil
}

func (in input) Floats() ([]float64, error) {
	var floatArr []float64 

	for _, v := range getLines(in.rawInput) {
		floatV, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}

		floatArr = append(floatArr, floatV)
	}

	return floatArr, nil
}
