package simplemath

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// LCM returns the least common multiple
// of two numbers
func LCM(a ...int) int {
	l := a[0]
	for i := 1; i < len(a); i++ {
		l = lcm(l, a[i])
	}
	return l
}

func lcm(a, b int) int {
	return a * b / GCD(a, b)
}

// GCD returns the greatest common denominator of
// a list of numbers
func GCD(a ...int) int {
	g := a[0]
	for i := 1; i < len(a); i++ {
		g = gcd(g, a[i])
	}
	return g
}

func gcd(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}
