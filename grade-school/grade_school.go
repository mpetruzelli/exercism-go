package school

import (
	"sort"
)

type Grade struct {
	i  int
	ss []string
}
type School map[int]*Grade

func insertSort(data []Grade, el Grade) []Grade {
	index := sort.Search(len(data), func(i int) bool { return data[i].i > el.i })
	data = append(data, Grade{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}

func New() *School {
	var sc School
	sc = make(School, 0)
	return &sc
}

func (s *School) Add(name string, i int) {
	if val, ok := (*s)[i]; ok {

		val.ss = append(val.ss, name)
		sort.Strings(val.ss)
	} else {

		var g Grade

		g = Grade{}

		g.i = i
		g.ss = make([]string, 0)
		g.ss = append(g.ss, name)
		(*s)[i] = &g
	}
}

func (s *School) Grade(i int) []string {
	if val, ok := (*s)[i]; ok {

		return val.ss

	}
	return []string{}
}

func (s *School) Enrollment() []Grade {
	v := make([]Grade, 0, len((*s)))

	for _, value := range *s {
		v = insertSort(v, *value)
	}
	return v
}
