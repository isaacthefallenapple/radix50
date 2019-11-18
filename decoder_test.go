package radix50

import "testing"

func TestDecoder(t *testing.T) {
	tests := map[string][]int16{
		"ABC":    {1683},
		"AB ":    {1680},
		"   ":    {0},
		" BB":    {82},
		"ABCABC": {1683, 1683},
		"  0":    {30},
	}

	for want, input := range tests {
		t.Run(want, func(t *testing.T) {
			dec := NewDecoder()
			if got := dec.Decode(input); got != want {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}
