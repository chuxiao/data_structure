package tree

import "testing"

func TestBinarySearchTree_Insert(t *testing.T) {
	list := []ElementType{10, 5, 7, 6, 15, 12, 13, 17, 16}
	expect := "5    6    7    10    12    13    15    16    17"
	bst := NewBinarySearchTree()
	for _, v := range list {
		bst.Insert(v)
	}
	s := bst.ToString()
	if s != expect {
		t.Errorf("output: %s, expect: %s\n", s, expect)
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	type test struct {
		input []ElementType
		delete []ElementType
		expect string
	}
	tt := []test {
		{
			input: []ElementType{},
			delete: []ElementType{10, 15},
			expect: "",
		},
		{
			input: []ElementType{10},
			delete: []ElementType{10},
			expect:"",
		},
		{
			input: []ElementType{10},
			delete: []ElementType{15},
			expect:"10",
		},
		{
			input: []ElementType{10, 11, 17, 15, 14},
			delete: []ElementType{15},
			expect:"10    11    14    17",
		},
		{
			input: []ElementType{10, 15, 17},
			delete: []ElementType{15},
			expect:"10    17",
		},
		{
			input: []ElementType{10, 15, 17, 16},
			delete: []ElementType{15},
			expect:"10    16    17",
		},
		{
			input: []ElementType{10, 15, 17, 18},
			delete: []ElementType{15},
			expect:"10    17    18",
		},
		{
			input: []ElementType{10, 15, 17, 16, 18},
			delete: []ElementType{15},
			expect:"10    16    17    18",
		},
	}
	for _, v := range tt {
		bst := NewBinarySearchTree()
		for _, v2 := range v.input {
			bst.Insert(v2)
		}
		for _, v2 := range v.delete {
			bst.Delete(v2)
		}
		s := bst.ToString()
		if s != v.expect {
			t.Errorf("input: %v, delete: %v, result: %s, expect: %s\n", v.input, v.delete, s, v.expect)
		}
	}

}

func TestBinarySearchTree_Find(t *testing.T) {

}

func TestBinarySearch(t *testing.T) {
	
}
