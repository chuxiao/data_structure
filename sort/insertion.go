package sort

func InsertionSort(arr []int) {
	size := len(arr)
	for i := 1; i < size; i++ {
		tmp := arr[i]
		j := i - 1
		for ;j >= 0; j-- {
			if tmp >= arr[j] {
				break
			}
			arr[j + 1] = arr[j]
		}
		arr[j + 1] = tmp
	}
}
