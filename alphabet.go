package radix50

func byteToInt(b byte) int16 {
	switch {
	case 'A' <= b && b <= 'Z':
		return int16(b) + 1 - 'A'
	case 'a' <= b && b <= 'z':
		return int16(b) + 1 - 'a'
	case b == '.':
		return 27
	case b == '?':
		return 28
	case b == '!':
		return 29
	case '0' <= b && b <= '9':
		return 30 + int16(b) - '0'
	default:
		return 0
	}
}

func intToByte(n int16) byte {
	switch {
	case 1 <= n && n <= 26:
		return byte(n - 1 + 'A')
	case 30 <= n && n <= 39:
		return byte(n - 30 + '0')
	case n == 27:
		return '.'
	case n == 28:
		return '?'
	case n == 29:
		return '!'
	default:
		return ' '
	}
}
