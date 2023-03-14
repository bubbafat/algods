package algods

import "testing"

func TestSet_Add_Unique_Adds(t *testing.T) {
	set := new(Set[int])
	for i := 0; i < 100; i++ {
		if set.Contains(i) {
			t.Error("Set should not already contain value", i)
		}

		set.Add(i)

		if !set.Contains(i) {
			t.Error("Set should contain value", i)
		}
	}
}

func TestSet_Add_Increments_Count(t *testing.T) {
	set := new(Set[int])
	for i := 0; i < 100; i++ {
		if count := set.Count(); count != i {
			t.Error("Count was not current before Add (expected, actual)", i, count)
		}

		set.Add(i)

		if count := set.Count(); count != i+1 {
			t.Error("Count was not current after Add (expected, actual)", i+1, count)
		}
	}
}

func TestSet_Add_Conflict_Does_Not_Add(t *testing.T) {
	count := 100
	set := new(Set[int])
	for i := 0; i < count; i++ {
		set.Add(i)
	}

	for i := 0; i < count; i++ {
		if set.Add(i) {
			t.Error("Adding a duplicate value should fail")
		}

		if actual := set.Count(); actual != count {
			t.Error("Count should not change when Add returns false (expected, actual)", count, actual)
		}
	}
}

func TestSet_Remove_Removes_Values(t *testing.T) {
	set := new(Set[int])
	for i := 0; i < 100; i++ {
		set.Add(i)
	}

	for i := 0; i < 100; i++ {
		if !set.Remove(i) {
			t.Error("Remove should have returned true", i)
		}

		if set.Contains(i) {
			t.Error("Contains should return false after Remove", i)
		}

		if count := set.Count(); count != 100-i-1 {
			t.Error("Count was not correct after remove (actual, expected)", count, 100-i-1)
		}
	}

	if count := set.Count(); count != 0 {
		t.Error("Count should be zero after all removals", count)
	}
}

func TestSet_Union_Unions(t *testing.T) {
	odds := new(Set[int])
	evens := new(Set[int])

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evens.Add(i)
		} else {
			odds.Add(i)
		}
	}

	all := odds.Union(evens)

	expectedCount := odds.Count() + evens.Count()

	if count := all.Count(); count != expectedCount {
		t.Error("Union count incorrect (actual, expected)", count, expectedCount)
	}

	for i := 0; i < 100; i++ {
		if !all.Contains(i) {
			t.Error("Union set should contain value", i)
		}
	}

	for i := 0; i < 100; i++ {
		all.Remove(i)
	}

	if count := all.Count(); count != 0 {
		t.Error("After removing each item the count should be zero", count)
	}
}

func TestSet_Intersection_None_Intersect(t *testing.T) {
	odds := new(Set[int])
	evens := new(Set[int])

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evens.Add(i)
		} else {
			odds.Add(i)
		}
	}

	all := odds.Intersection(evens)

	if count := all.Count(); count != 0 {
		t.Error("Count should be 0 for non-interesting sets")
	}
}

func TestSet_Intersection_All_Intersect(t *testing.T) {
	odds := new(Set[int])
	evens := new(Set[int])

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evens.Add(i)
		} else {
			odds.Add(i)
		}
	}

	onlyOdds := odds.Intersection(odds)

	if count := onlyOdds.Count(); count != odds.Count() {
		t.Error("Complete intersection should result in the same output set")
	}

	odds.ForEach(func(value int) bool {
		return onlyOdds.Remove(value)
	})

	if count := onlyOdds.Count(); count != 0 {
		t.Error("All values should have been removed from the set", count)
	}
}
