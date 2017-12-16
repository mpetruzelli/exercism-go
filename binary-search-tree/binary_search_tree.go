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

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func (t *SearchTreeData) Walk(ch chan int) {
	if t == nil {
		return
	}
	t.left.Walk(ch)
	ch <- t.data
	t.right.Walk(ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func (t *SearchTreeData) Walker() <-chan int {
	ch := make(chan int)
	go func() {
		t.Walk(ch)
		close(ch)
	}()
	return ch
}

func Bst(v int) SearchTreeData {
	var t SearchTreeData
	t = SearchTreeData{nil, v, nil}
	return t
}

func (t *SearchTreeData) MapString(x func(int) string) (out []string) {
	c1 := t.Walker()
	for {
		v1, ok := <-c1
		if !ok {
			break
		}
		out = append(out, x(v1))
	}
	return
}

func (t *SearchTreeData) MapInt(x func(int) int) (out []int) {
	c1 := t.Walker()
	for {
		v1, ok := <-c1
		if !ok {
			break
		}
		out = append(out, x(v1))
	}

	return
}
