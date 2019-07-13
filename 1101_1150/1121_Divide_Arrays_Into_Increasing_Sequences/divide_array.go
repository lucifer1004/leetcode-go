package dividearray

func canDivideIntoSubsequences(nums []int, K int) bool {
	// Count frequency of each number
	freq := []int{0}
	curr := nums[0]
	idx := 0
	for _, v := range nums {
		if v == curr {
			freq[idx]++
		} else {
			freq = append(freq, 1)
			idx++
			curr = v
		}
	}

	// Find the maximal frequency
	maxFreq := 0
	for _, v := range freq {
		if v > maxFreq {
			maxFreq = v
		}
	}

	// To produce increasing sequences, the numbers should be divided into at least `maxFreq` groups, thus determining the maximum lower bound of sequence length.
	return K <= len(nums)/maxFreq
}
