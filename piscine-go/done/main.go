package main

import "os"

const (
	maxInt64 = 1<<63 - 1
	minInt64 = -1 << 63
)

func parseInt64(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	i := 0
	neg := false
	if s[0] == '+' {
		i++
	} else if s[0] == '-' {
		neg = true
		i++
	}
	if i == len(s) {
		return 0, false
	}
	var acc uint64 = 0
	var limit uint64
	if neg {
		limit = uint64(1) << 63
	} else {
		limit = uint64(maxInt64)
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := uint64(c - '0')
		if acc > (limit-d)/10 {
			return 0, false
		}
		acc = acc*10 + d
	}
	if neg {
		if acc == (uint64(1) << 63) {
			return minInt64, true
		}
		return -int64(acc), true
	}
	return int64(acc), true
}

func itoaInt64(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	var u uint64
	if neg {
		u = uint64(-(n + 1)) + 1
	} else {
		u = uint64(n)
	}
	var buf [20]byte
	i := len(buf)
	for u > 0 {
		i--
		buf[i] = byte('0' + (u % 10))
		u /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}

func writeString(s string) {
	os.Stdout.Write([]byte(s))
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, ok1 := parseInt64(aStr)
	b, ok2 := parseInt64(bStr)
	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		if (b > 0 && a > maxInt64-b) || (b < 0 && a < minInt64-b) {
			return
		}
		writeString(itoaInt64(a+b) + "\n")
	case "-":
		if (b < 0 && a > maxInt64+b) || (b > 0 && a < minInt64+b) {
			return
		}
		writeString(itoaInt64(a-b) + "\n")
	case "*":
		temp := a * b
		if a != 0 && temp/a != b {
			return
		}
		writeString(itoaInt64(temp) + "\n")
	case "/":
		if b == 0 {
			writeString("No division by 0\n")
			return
		}
		writeString(itoaInt64(a/b) + "\n")
	case "%":
		if b == 0 {
			writeString("No modulo by 0\n")
			return
		}
		writeString(itoaInt64(a%b) + "\n")
	default:
		return
	}
}
