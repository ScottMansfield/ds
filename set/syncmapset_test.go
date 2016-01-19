package set_test

import "testing"

import set "."

func TestNewSyncMapSet(t *testing.T) {
    set.NewSyncMapSet()
}

func BenchmarkNewSyncMapSet(b *testing.B) {
    for i := 0; i < b.N; i++ {
        set.NewSyncMapSet()
    }
}

func TestSyncMapAdd(t *testing.T) {
    s := set.NewSyncMapSet()
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

func BenchmarkSyncMapAdd(b *testing.B) {
    s := set.NewSyncMapSet()
    for i := 0; i < b.N; i++ {
        if err := s.Add(i); err != nil {
            b.Fail()
        }
    }
}

func TestSyncMapAddAll(t *testing.T) {
    s := set.NewSyncMapSet()
    data := genKeyData(10)
    if err := s.AddAll(data...); err != nil {
        t.Fail()
    }
    if sz, _ := s.Size(); sz != 10 {
        t.Fail()
    }
}

func BenchmarkSyncMapAddAll(b *testing.B) {
    s := set.NewSyncMapSet()
    data := genKeyData(1000)
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        if err := s.AddAll(data...); err != nil {
            b.Fail()
        }
    }
}

func TestSyncMapContains(t *testing.T) {
    s := set.NewSyncMapSet()
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

func BenchmarkSyncMapContains(b *testing.B) {
    s := set.NewSyncMapSet()
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

func TestSyncMapContainsAll(t *testing.T) {
    s := set.NewSyncMapSet()
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

func BenchmarkSyncMapContainsAll(b *testing.B) {
    s := set.NewSyncMapSet()
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

func TestSyncMapRemove(t *testing.T) {
    s := set.NewSyncMapSet()

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
func BenchmarkSyncMapRemove(b *testing.B) {
    sets := make([]set.Set, b.N)

    for i := 0; i < b.N; i++ {
        sets[i] = set.NewSyncMapSet()
        sets[i].Add(4)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        if err := sets[i].Remove(4); err != nil {
            b.Fail()
        }
    }
}

func TestSyncMapRemoveAll(t *testing.T) {
    s := set.NewSyncMapSet()
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

func BenchmarkSyncMapRemoveAll(b *testing.B) {
    sets := make([]set.Set, b.N)
    keys := make([][]interface{}, b.N)
    data := genKeyData(100)

    for i := 0; i < b.N; i++ {
        sets[i] = set.NewSyncMapSet()
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

func TestSyncMapClear(t *testing.T) {
    s := set.NewSyncMapSet()
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

func BenchmarkSyncMapClear(b *testing.B) {
    sets := make([]set.Set, 1000)
    for i := 0; i < len(sets); i++ {
        sets[i] = set.NewSyncMapSet()
    }
    b.ResetTimer()

    for i, j := 0, 0; i < b.N; i, j = i+1, j+1 {
        if i % 1000 == 0 {
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

func TestSyncMapSize(t *testing.T) {
    s := set.NewSyncMapSet()
    s.Add(4)
    size, err := s.Size()
    if err != nil {
        t.Fail()
    }
    if size != 1 {
        t.Fail()
    }
}

func BenchmarkSyncMapSize(b *testing.B) {
    s := set.NewSyncMapSet()
    s.AddAll(genKeyData(1000)...)
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        if _, err := s.Size(); err != nil {
            b.Fail()
        }
    }
}

func TestSyncMapIsEmpty(t *testing.T) {
    s := set.NewSyncMapSet()

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

func BenchmarkSyncMapIsEmpty(b *testing.B) {
    s1 := set.NewSyncMapSet()
    s2 := set.NewSyncMapSet()
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

func TestSyncMapToSlice(t *testing.T) {
    s := set.NewSyncMapSet()
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

func BenchmarkSyncMapToSlice(b *testing.B) {
    s := set.NewSyncMapSet()
    s.AddAll(genKeyData(100)...)
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        if _, err := s.ToSlice(); err != nil {
            b.Fail()
        }
    }
}
