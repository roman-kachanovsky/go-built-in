package functools

import "testing"

var float64Cases = []struct {
	in interface{}
	want float64
	e bool
}{
	{0, 0.0, false},
	{1, 1.0, false},
	{-1, -1.0, false},
	{0.0, 0.0, false},
	{0.1, 0.1, false},
	{-0.1, -0.1, false},
	{uint(2), 2.0, false},
	{uint8(0), 0.0, false},
	{uint16(1), 1.0, false},
	{uint32(0), 0.0, false},
	{uint64(5), 5.0, false},
	{int8(0), 0.0, false},
	{int16(1), 1.0, false},
	{int32(0), 0.0, false},
	{int64(5), 5.0, false},
	{"abc", 0.0, true},
	{"", 0.0, true},
	{[]int{}, 0.0, true},
	{[]int(nil), 0.0, true},
	{[]int{1, 2, 3}, 0.0, true},
	{[3]int{}, 0.0, true},
	{[]float64{}, 0.0, true},
	{[]float64{1.0, 2.1, 3.5}, 0.0, true},
	{[]string{}, 0.0, true},
	{[]string{"a", "b", "c"}, 0.0, true},
	{[]rune{}, 0.0, true},
	{[]rune{'a', 'b', 'c'}, 0.0, true},
	{float32(0.001), 0.0010000000474974513, false},
	{nil, 0.0, false},
	{'a', 97.0, false},
	{byte('a'), 97.0, false},
	{make(map[string]int), 0.0, true},
	{map[string]int{"a": 1}, 0.0, true},
	{true, 1.0, false},
	{false, 0.0, false},
}

func TestToFloat64Safe(t *testing.T) {
	for _, c := range float64Cases {
		got, err := ToFloat64Safe(c.in)

		if err != nil && !c.e {
			t.Errorf("ToFloat64Safe(%v) raised %v", c.in, err)
		}

		if err == nil && got != c.want {
			t.Errorf("ToFloat64Safe(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestToFloat64(t *testing.T) {
	for _, c := range float64Cases {
		if !c.e {
			got := ToFloat64(c.in)

			if got != c.want {
				t.Errorf("ToFloat64(%v) == %v want %v", c.in, got, c.want)
			}
		}
	}
}
