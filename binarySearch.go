package algorithm

func search(arr []int, des int, min int, max int) int {
	idx := min + (max-min)/2
	if (idx == min || idx == max) && arr[idx] != des {
		return -1
	}
	if arr[idx] == des {
		return idx
	} else if arr[idx] > des {
		return search(arr, des, min, idx)
	} else {
		return search(arr, des, idx, max)
	}
}

func Search(arr []int, des int) int {
	return search(arr, des, 0, len(arr))
}
