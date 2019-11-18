package radix50

import (
	"reflect"
	"testing"
)

func TestEncoder(t *testing.T) {
	tests := map[string][]int16{
		"abc":    {1683},
		"   ":    {0},
		"abcabc": {1683, 1683},
		"abcab":  {1683, 1680},
		"":       {0},
		" aa":    {41},
		"  0":    {30},
	}

	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			enc := NewEncoder()
			if got := enc.Encode(input); !reflect.DeepEqual(want, got) {
				t.Errorf("want %#v, got %#v", want, got)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	tests := map[string]struct {
		toWrite []byte
		want    []byte
	}{
		"normal input": {
			toWrite: []byte("hello world"),
			want:    []byte("hello world\x00"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			enc := NewEncoder()
			enc.Write(test.toWrite)
			if !reflect.DeepEqual(enc.buf, test.want) {
				t.Errorf("got %#v, want %#v", enc.buf, test.want)
			}
		})
	}
}

func TestRead(t *testing.T) {
	tests := map[string][]byte{
		"ABC":    {0b00000110, 0b10010011},
		"AB ":    {0b110, 0b10010000},
		"     A": {0, 0, 0, 1},
	}

	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			enc := NewEncoder()
			enc.Write([]byte(input))
			p := make([]byte, len(enc.buf)/3*2)
			enc.Read(p)
			if !reflect.DeepEqual(p, want) {
				t.Errorf("want %#v, got %#v", want, p)
			}
		})
	}
}

func TestMultByShifting(t *testing.T) {
	for c1 := range [40]struct{}{} {
		for c2 := range [40]struct{}{} {
			for c3 := range [40]struct{}{} {
				if 40*40*c1+40*c2+c3 != (c1<<10+c1<<9+c1<<6)+(c2<<5+c2<<3)+c3 {
					t.Errorf("wrong for for c1=%d, c2=%d, c3=%d", c1, c2, c3)
				}
			}
		}
	}
}
