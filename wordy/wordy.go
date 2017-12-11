package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(t *Tree, ch chan string) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *Tree) <-chan string {
	ch := make(chan string)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}
func insert(t *Tree, lr, v string) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	switch lr {
	case "left":
		t.Left = insert(t.Left, "right", v)
		return t
	case "right":
		t.Right = insert(t.Right, "left", v)
		return t
	}
	return t
}

var operation = map[string]bool{
	"plus":       true,
	"minus":      true,
	"multiplied": true,
	"divided":    true,
	"raised":     true,
}

//creates new binary tree inserting left,right,left,right
func New(split []string) *Tree {
	var t *Tree
	var order string = "right"
	for _, v := range split {
		t = insert(t, order, v)
		if order == "left" {
			order = "right"
		} else {
			order = "left"
		}
	}
	return t
}

func Answer(s string) (int, bool) {
	r := strings.Replace(s, "?", "", -1)
	t := strings.Fields(r)
	result := 0
	numbersandoperations := []string{}

	for _, v := range t {
		if _, err := strconv.Atoi(v); err == nil {
			numbersandoperations = append(numbersandoperations, v)
		}
		if _, ok := operation[v]; ok == true {
			numbersandoperations = append(numbersandoperations, v)
		}
	}

	if len(numbersandoperations) < 3 {
		return result, false
	}

	tree := New(numbersandoperations)

	ch := Walker(tree)
	for {
		x, ok := <-ch
		fmt.Println("ch1", x, ok)
		if !ok {
			ch = nil
		}
		if ch == nil {
			break
		}
	}
	return result, true
}
