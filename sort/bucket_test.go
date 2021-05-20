package sort

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	a := []int{64, 8, 216, 512, 27, 729, 0, 1, 343, 125}
	BucketSort(a, 3)
	fmt.Println(a)
}
