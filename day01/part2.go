package day01

// Compute the similarity score of elements of a with respect to b.
// The similarity score is the addition of each number in a after
// multiplying it by the number of times that it appears in b.
func SimilarityScore(a, b []int) int {
	return computeScore(a, buildCounter(b), 0)
}

func buildCounter(numbers []int) map[int]int {
	counter := make(map[int]int)
	for _, num := range numbers {
		count, ok := counter[num]
		if !ok {
			counter[num] = 1
		} else {
			counter[num] = count + 1
		}
	}
	return counter
}

func computeScore(a []int, counter map[int]int, acc int) int {
	if len(a) == 0 {
		return acc
	}
	count, ok := counter[a[0]]
	if !ok {
		return computeScore(a[1:], counter, acc)
	}
	return computeScore(a[1:], counter, acc+(a[0]*count))

}
