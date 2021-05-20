package sort

type bucket [10][]int

func (b *bucket) add(d int, v int) {
	b[d] = append(b[d], v)
}

func (b *bucket) reset() {
	for i := 0; i < 10; i++ {
		if b[i] != nil {
			b[i] = b[i][:0]
		}
	}
}

func BucketSort(array []int, loop int) {
	var b bucket
	base := 1
	for i := 0; i < loop; i++ {
		base *= 10
		for _, v := range array {
			d := getDigital(v, base)
			b.add(d, v)
		}
		array = array[:0]
		for j := 0; j < 10; j++ {
			if len(b[j]) == 0 {
				continue
			}
			array = append(array, b[j]...)
		}
		b.reset()
	}

}

func getDigital(v, base int) int {
	return v / (base / 10) - v / base * 10
}