package radix50

type Decoder struct{}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (dec *Decoder) Decode(input []int16) string {
	out := make([]byte, len(input)*3)
	var p int
	for _, n := range input {
		c3, c2, c1 := n%40, (n/40)%40, (n/40/40)%40
		out[p], out[p+1], out[p+2] = intToByte(c1), intToByte(c2), intToByte(c3)
		p += 3
	}
	return string(out)
}
