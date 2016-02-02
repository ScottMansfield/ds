package set_test

import "testing"

import set "."

func TestNewSortedSet(t *testing.T) {
	set.NewSortedSet()
}

func BenchmarkNewSortedSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set.NewSortedSet()
	}
}

func TestSortedAdd(t *testing.T) {
	s := set.NewSortedSet()
	if err := s.Add(4); err != nil {
		t.Fail()
	}
	// Assume other methods not under test work
	if sz, _ := s.Size(); sz != 1 {
		t.Fail()
	}

	// Repeat to test set property of uniqueness
	if err := s.Add(4); err != nil {
		t.Fail()
	}
	// Assume other methods not under test work
	if sz, _ := s.Size(); sz != 1 {
		t.Fail()
	}
}

func BenchmarkSortedAdd(b *testing.B) {
	s := set.NewSortedSet()
	for i := 0; i < b.N; i++ {
		if err := s.Add(i); err != nil {
			b.Fail()
		}
	}
}

func TestSortedAddAll(t *testing.T) {
	s := set.NewSortedSet()
	data := genKeyData(10)
	if err := s.AddAll(data...); err != nil {
		t.Fail()
	}
	if sz, _ := s.Size(); sz != 10 {
		t.Fail()
	}
}

func BenchmarkSortedAddAll(b *testing.B) {
	s := set.NewSortedSet()
	data := genKeyData(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := s.AddAll(data...); err != nil {
			b.Fail()
		}
	}
}

func TestSortedContains(t *testing.T) {
	s := set.NewSortedSet()
	s.Add(4)

	c, err := s.Contains(4)
	if !c || err != nil {
		t.Fail()
	}

	c, err = s.Contains(5)
	if c || err != nil {
		t.Fail()
	}
}

func BenchmarkSortedContains(b *testing.B) {
	s := set.NewSortedSet()
	s.Add(4)

	for i := 0; i < b.N/2; i++ {
		c, err := s.Contains(4)
		if !c || err != nil {
			b.Fail()
		}
	}

	for i := 0; i < b.N/2; i++ {
		c, err := s.Contains(5)
		if c || err != nil {
			b.Fail()
		}
	}
}

func TestSortedContainsAll(t *testing.T) {
	s := set.NewSortedSet()
	data := genKeyData(10)

	s.AddAll(data...)

	c, err := s.ContainsAll(data...)
	if !c || err != nil {
		t.Fatalf("")
	}

	// Add a known missing key
	data = append(data, -1)

	c, err = s.ContainsAll(data...)
	if c || err != nil {
		t.Fail()
	}
}

func BenchmarkSortedContainsAll(b *testing.B) {
	s := set.NewSortedSet()
	data := genKeyData(1000)
	s.AddAll(data...)
	b.ResetTimer()

	for i := 0; i < b.N/2; i++ {
		c, err := s.ContainsAll(data...)
		if !c || err != nil {
			b.Fail()
		}
	}

	// Refresh to get a mix of hits and misses
	b.StopTimer()
	data = genKeyData(1000)
	b.StartTimer()

	for i := 0; i < b.N/2; i++ {
		_, err := s.ContainsAll(data...)
		if err != nil {
			b.Fail()
		}
	}
}

func TestSortedRemove(t *testing.T) {
	s := set.NewSortedSet()

	// Remove something that doesn't exist
	if err := s.Remove(4); err != nil {
		t.Fail()
	}
	if sz, _ := s.Size(); sz != 0 {
		t.Fail()
	}

	// Remove something that does exist
	s.Add(4)
	if err := s.Remove(4); err != nil {
		t.Fail()
	}
	if sz, _ := s.Size(); sz != 0 {
		t.Fail()
	}
}

// Not perfect because different implementations might
// be slower the larger it gets
func BenchmarkSortedRemove(b *testing.B) {
	sets := make([]set.Set, b.N)

	for i := 0; i < b.N; i++ {
		sets[i] = set.NewSortedSet()
		sets[i].Add(4)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := sets[i].Remove(4); err != nil {
			b.Fail()
		}
	}
}

func TestSortedRemoveAll(t *testing.T) {
	s := set.NewSortedSet()
	data := genKeyData(1000)
	s.AddAll(data...)

	if err := s.RemoveAll(data...); err != nil {
		t.Fail()
	}
	if sz, _ := s.Size(); sz != 0 {
		t.Fail()
	}

	// Refresh to get a mix of hits and misses
	oldData := data
	data = genKeyData(1000)
	s.AddAll(oldData...)

	if err := s.RemoveAll(data...); err != nil {
		t.Fail()
	}
}

func BenchmarkSortedRemoveAll(b *testing.B) {
	sets := make([]set.Set, b.N)
	keys := make([][]interface{}, b.N)
	data := genKeyData(100)

	for i := 0; i < b.N; i++ {
		sets[i] = set.NewSortedSet()
		sets[i].AddAll(data...)
		keys[i] = genKeyData(100)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := sets[i].RemoveAll(keys[i]...); err != nil {
			b.Fail()
		}
	}
}

func TestSortedClear(t *testing.T) {
	s := set.NewSortedSet()
	s.AddAll(genKeyData(1000)...)

	if err := s.Clear(); err != nil {
		t.Fail()
	}
	if sz, _ := s.Size(); sz != 0 {
		t.Fail()
	}

	// repeat in case Clear broke something
	s.AddAll(genKeyData(1000)...)

	if err := s.Clear(); err != nil {
		t.Fail()
	}
	if sz, _ := s.Size(); sz != 0 {
		t.Fail()
	}
}

func BenchmarkSortedClear(b *testing.B) {
	sets := make([]set.Set, 1000)
	for i := 0; i < len(sets); i++ {
		sets[i] = set.NewSortedSet()
	}
	b.ResetTimer()

	for i, j := 0, 0; i < b.N; i, j = i+1, j+1 {
		if i%1000 == 0 {
			b.StopTimer()
			fillSets(sets)
			j = 0
			b.StartTimer()
		}
		if err := sets[j].Clear(); err != nil {
			b.Fail()
		}
	}
}

func TestSortedSize(t *testing.T) {
	s := set.NewSortedSet()
	s.Add(4)
	size, err := s.Size()
	if err != nil {
		t.Fail()
	}
	if size != 1 {
		t.Fail()
	}
}

func BenchmarkSortedSize(b *testing.B) {
	s := set.NewSortedSet()
	s.AddAll(genKeyData(1000)...)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := s.Size(); err != nil {
			b.Fail()
		}
	}
}

func TestSortedIsEmpty(t *testing.T) {
	s := set.NewSortedSet()

	empty, err := s.IsEmpty()
	if err != nil {
		t.Fail()
	}
	if !empty {
		t.Fail()
	}

	s.Add(4)

	empty, err = s.IsEmpty()
	if err != nil {
		t.Fail()
	}
	if empty {
		t.Fail()
	}
}

func BenchmarkSortedIsEmpty(b *testing.B) {
	s1 := set.NewSortedSet()
	s2 := set.NewSortedSet()
	s2.AddAll(genKeyData(1000)...)
	b.ResetTimer()

	for i := 0; i < b.N/2; i++ {
		empty, err := s1.IsEmpty()
		if err != nil {
			b.Fail()
		}
		if !empty {
			b.Fail()
		}
	}

	for i := 0; i < b.N/2; i++ {
		empty, err := s2.IsEmpty()
		if err != nil {
			b.Fail()
		}
		if empty {
			b.Fail()
		}
	}
}

func TestSortedToSlice(t *testing.T) {
	s := set.NewSortedSet()
	s.Add(4)

	slice, err := s.ToSlice()
	if err != nil {
		t.Fail()
	}
	if len(slice) != 1 {
		t.Fail()
	}

	s.Add(5)

	slice, err = s.ToSlice()
	if err != nil {
		t.Fail()
	}
	if len(slice) != 2 {
		t.Fail()
	}
}

func BenchmarkSortedToSlice(b *testing.B) {
	s := set.NewSortedSet()
	s.AddAll(genKeyData(100)...)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := s.ToSlice(); err != nil {
			b.Fail()
		}
	}
}
