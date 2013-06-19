package algorithm

func Search(arr []int, des int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	middle, low, high := 0, 0, length-1
	if arr[low] == des {
		return low
	}
	if arr[high] == des {
		return high
	}
	for low <= high {
		middle = low + (high-low)/2
		if arr[middle] == des {
			return middle
		}
		if arr[middle] > des {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}
