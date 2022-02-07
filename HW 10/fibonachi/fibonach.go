package fibonachi

func FibonachiProcess(x float64) (float64, error) {
	if x == 0 {
		return 0.0, nil
	}

	if x == 1.0 {
		return 1.0, nil
	}

	result1, err := FibonachiProcess(x - 1)
	result2, err := FibonachiProcess(x - 2)

	return result1 + result2, err
}
