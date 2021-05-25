package sort

func ShellSort(arr []int) {
	size := len(arr)
	for inc := size / 2; inc > 0; inc /= 2 {
		for i := inc; i < size; i += inc {
			tmp := arr[i]
			j := i - inc
			for ; j >= 0; j -= inc {
				if tmp >= arr[j] {
					break
				}
				arr[j + inc] = arr[j]
			}
			arr[j + inc] = tmp
		}
	}
}