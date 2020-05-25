package public

import "sort"

type Base interface {
	Less(Base) bool
}

type MapSorter struct {
	Keys []string
	Vals []Base
}

func NewMapSorter(m map[string]Base) *MapSorter {
	ms := &MapSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]Base, 0, len(m)),
	}
	for k, v := range m {
		ms.Keys = append(ms.Keys, k)
		ms.Vals = append(ms.Vals, v)
	}
	return ms
}

// Sort sort tasker map
func (ms *MapSorter) Sort() {
	sort.Sort(ms)
}

func (ms *MapSorter) Len() int {
	return len(ms.Keys)
}

func (ms *MapSorter) Less(i, j int) bool {
	if nil == ms.Vals[i] {
		return false
	}
	if nil == ms.Vals[j] {
		return true
	}
	return ms.Vals[i].Less(ms.Vals[j])
}

func (ms *MapSorter) Swap(i, j int) {
	ms.Vals[i], ms.Vals[j] = ms.Vals[j], ms.Vals[i]
	ms.Keys[i], ms.Keys[j] = ms.Keys[j], ms.Keys[i]
}
