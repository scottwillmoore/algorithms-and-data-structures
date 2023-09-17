package list

import "testing"

func TestArrayList(t *testing.T) {
	list := New[rune]()
	list.Add(0, 'A')
	if list.Get(0) != 'A' {
		t.Errorf("")
	}
}
