package bag

import "testing"

func TestBag_Add(t *testing.T) {
	bag := New()
	for i := 0; i < 10; i++ {
		bag.Add(i)
	}

	if s := bag.Size(); s != 10 {
		t.Errorf("Size error: excepted %v, got %v", 10, s)
	}

	if e := bag.Empty(); e {
		t.Errorf("Empty error: excepted %v, got %v", false, e)
	}
}

func TestBag_Iterator(t *testing.T) {
	bag := New()
	iter1 := bag.Iter()

	if i := iter1.HasNext(); i {
		t.Errorf("HasNext Error: excepted %v, got %v", false, i)
	}

	for i := 0; i < 10; i++ {
		bag.Add(i)
	}

	iter2 := bag.Iter()

	j := 9
	for iter2.HasNext() {
		if item := iter2.Value(); item != j {
			t.Errorf("Next Error: excepted %v, got %v", j, item)
		}
		j--
	}
}
