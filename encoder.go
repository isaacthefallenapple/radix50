package radix50

import (
	"io"
)

type Encoder struct {
	buf []byte
}

func NewEncoder() *Encoder {
	return &Encoder{
		buf: make([]byte, 0),
	}
}

func (enc *Encoder) Write(p []byte) (n int, err error) {
	buf := enc.buf
	buflen, plen := len(buf), len(p)
	if plen == 0 {
		return 0, nil
	}
	newlen := buflen + plen
	newlen += (3 - (newlen % 3)) % 3
	enc.buf = make([]byte, newlen)
	n = copy(enc.buf, append(buf, p...)) - buflen
	return
}

func (enc *Encoder) Read(p []byte) (n int, err error) {
	encoded := enc.Encode(string(enc.buf))
	var i int
	for _, b := range encoded {
		hi, lo := b>>8, b&0xff
		plen := len(p)
		if i < plen {
			p[i] = byte(hi)
		} else {
			break
		}
		i++
		if i < plen {
			p[i] = byte(lo)
		} else {
			break
		}
		i++
	}
	n = i
	if n < len(p) {
		err = io.EOF
	}
	return
}

func (enc *Encoder) Encode(s string) []int16 {
	strlen := len(s)
	if strlen == 0 {
		return []int16{0}
	}
	padded := make([]byte, strlen)
	copy(padded, s)
	for r := len(padded) % 3; r > 0; r = len(padded) % 3 {
		padded = append(padded, ' ')
	}
	padlen := len(padded)
	out := make([]int16, padlen/3)
	var i, p int

	for ; i < padlen; i += 3 {
		c1, c2, c3 := byteToInt(padded[i]), byteToInt(padded[i+1]), byteToInt(padded[i+2])
		out[p] = (c1<<10 + c1<<9 + c1<<6) + (c2<<5 + c2<<3) + c3
		p++
	}
	return out
}
