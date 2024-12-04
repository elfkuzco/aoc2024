package day02

// Check if a report is safe with the dampener.
// If a report is originally unsafe, the dampener tries to see if it can be
// made safe by removing a single element.
func IsSafeWithDampener(report []int) bool {
	if IsSafe(report) {
		return true
	}
	return isSafeWithoutMember(report, 0)
}

func isSafeWithoutMember(report []int, index int) bool {
	if index == len(report) {
		return false
	}
	reportCopy := make([]int, len(report))
	copy(reportCopy, report)
	// check if it's safe by removing the element at index
	reportCopy = append(reportCopy[:index], reportCopy[index+1:]...)
	if IsSafe(reportCopy) {
		return true
	}
	return isSafeWithoutMember(report, index+1)
}
