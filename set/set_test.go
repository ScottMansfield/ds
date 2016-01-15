package set

import "math/rand"
import "testing"
import "time"

// utility functions
var r *rand.Rand
func init() {
    r = rand.New(rand.NewSource(time.Now().Unix()))
}

func genKeyData(num int) []interface{} {
    ret := make([]interface{}, num)

    for i := 0; i < num; i++ {
        ret[i] = r.Intn(10000)
    }

    return ret
}

// Tests

func TestNewSet(t *testing.T) {
    New()
}

func BenchmarkNewSet(b *testing.B) {
    for i := 0; i < b.N; i++ {
        New()
    }
}

func TestAdd(t *testing.T) {
    s := New()
    if err := s.Add(4); err != nil {
        t.Fail()
    }
    if sz, _ := s.Size(); sz != 1 {
        t.Fail()
    }
}

func BenchmarkAdd(b *testing.B) {
    s := New()
    for i := 0; i < b.N; i++ {
        if err := s.Add(i); err != nil {
            b.Fail()
        }
    }
}

func TestAddAll(t *testing.T) {
    s := New()
    data := genKeyData(10)
    if err := s.AddAll(data...); err != nil {
        t.Fail()
    }
    if sz, _ := s.Size(); sz != 10 {
        t.Fail()
    }
}

func BenchmarkAddAll(b *testing.B) {
    s := New()
    for i := 0; i < b.N; i++ {
        b.StopTimer()
        data := genKeyData(1000)
        b.StartTimer()

        if err := s.AddAll(data...); err != nil {
            b.Fail()
        }
    }
}

func TestContains(t *testing.T) {
    s := New()
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

func BenchmarkContains(b *testing.B) {
    s := New()
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

func TestContainsAll(t *testing.T) {
    s := New()
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

func BenchmarkContainsAll(b *testing.B) {
    s := New()
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

func TestRemove(t *testing.T) {
    s := New()

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

func BenchmarkRemove(b *testing.B) {
    sets := make([]Set, b.N)

    for i := 0; i < b.N; i++ {
        sets[i] = New()
        sets[i].Add(4)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        if err := sets[i].Remove(4); err != nil {
            b.Fail()
        }
    }
}

func TestRemoveAll(t *testing.T) {
    s := New()
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

func BenchmarkRemoveAll(b *testing.B) {
    sets := make([]Set, b.N)
    keys := make([][]interface{}, b.N)

    for i := 0; i < b.N; i++ {
        sets[i] = New()
        sets[i].AddAll(genKeyData(1000)...)
        keys[i] = genKeyData(1000)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        if err := sets[i].RemoveAll(keys[i]...); err != nil {
            b.Fail()
        }
    }
}

func TestClear(t *testing.T) {
    s := New()
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

func BenchmarkClear(b *testing.B) {
    println(b.N)
    sets := make([]Set, b.N)

    for i := 0; i < b.N; i++ {
        sets[i] = New()
        sets[i].AddAll(genKeyData(100)...)
    }

    println("resetting timer")
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        if err := sets[i].Clear(); err != nil {
            b.Fail()
        }
    }
}

//size

//isEmpty

//toslice
