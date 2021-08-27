package check

func Checking(scores [8]int) (avg float64, filtered []int) {
	avg = average(scores)
	filtered = filter(scores)

	return
}

func average(scores [8]int) (hasil float64) {
	var sum int
	for _, score := range scores {
		sum += score
	}

	hasil = float64(sum) / float64(len(scores))
	return
}

func filter(scores [8]int) (hasil []int) {
	for _, score := range scores {
		if score > 90 {
			hasil = append(hasil, score)
		}
	}
	return
}
