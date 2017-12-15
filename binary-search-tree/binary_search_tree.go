package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func (t *SearchTreeData) Insert(v int) {

	switch {
	case v <= t.data:
		if t.left != nil {
			t.left.Insert(v)
		} else {
			t.left = &SearchTreeData{nil, v, nil}
		}
		break
	case v > t.data:
		if t.right != nil {
			t.right.Insert(v)
		} else {
			t.right = &SearchTreeData{nil, v, nil}
		}
		break
	}

}

func Bst(v int) SearchTreeData {
	var t SearchTreeData
	t = SearchTreeData{nil, v, nil}
	return t
}

func (*SearchTreeData) MapString(func(int) string) []string {
	return []string{}
}

func (*SearchTreeData) MapInt(func(int) int) []int {
	return []int{}
}
