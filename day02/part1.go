package day02

// Check if entries of report are safe. A report is safe if:
// - the entries are either all increasing or all decreasing
// - any two adjacent levels differ by at least one and at most three.
func IsSafe(report []int) bool {
	return differencesAreSafe(runningDifference(report), 0)
}

// Get the slice of the running difference of the elements of slice.
// The running difference is the diference beteen an element and the element
// before it
func runningDifference(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	return computeRunningDifference(a, []int{}, 0)
}

func computeRunningDifference(a []int, acc []int, index uint) []int {
	// If index is 0, proceed to compute the running difference from
	// the index at 1.
	if index == 0 {
		return computeRunningDifference(a, acc, 1)
	}
	// if index is equal to the length, then there is no entry to be computed
	// next, return the accumulated result.
	if int(index) == len(a) {
		return acc
	}
	// add the difference and increment the index
	acc = append(acc, a[index-1]-a[index])
	return computeRunningDifference(a, acc, index+1)
}

// Check if two non-zero integers have the same sign
func signsAreTheSame(a, b int) bool {
	sum := a + b
	return (sum > a && sum > b) || (sum < a && sum < b)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return a * -1
}

func differencesAreSafe(differences []int, index uint) bool {
	if index == 0 {
		if differences[index] == 0 || abs(differences[index]) > 3 {
			return false
		}
		return differencesAreSafe(differences, 1)
	}

	// if index is equal to the length, then there is no more work to be done
	if int(index) == len(differences) {
		return true
	}

	if differences[index] == 0 || !signsAreTheSame(differences[index], differences[index-1]) || abs(differences[index]) > 3 {
		return false
	}
	return differencesAreSafe(differences, index+1)
}

// Check if two slices are equal
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
