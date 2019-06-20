package searchrange

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length == 0 || target < nums[0] || target > nums[length-1] {
		return []int{-1, -1}
	}

	left := -1
	right := -1

	l := 0
	r := length - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			if mid == 0 {
				left = 0
				break
			}
			if nums[mid-1] < target {
				left = mid
				break
			}
		}
		if nums[mid] > target || (mid > 0 && nums[mid-1] >= target) {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}

	l = 0
	r = length - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			if mid == length-1 {
				right = length - 1
				break
			}
			if nums[mid+1] > target {
				right = mid
				break
			}
		}
		if nums[mid] < target || (mid < length-1 && nums[mid+1] <= target) {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}

	return []int{left, right}
}
