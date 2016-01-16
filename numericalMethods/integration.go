package numericalMethods

// Euler1D is for solving the numerical integration 1D euler method
func Euler1D(a float64, b float64, N int, initValue float64, f func(float64, float64) float64) float64 {
	h := (b - a) / float64(N)
	t := a
	omega := initValue

	for i := 0; i < N; i++ {
		omega += h * f(t, omega)
		t += h
	}

	return omega
}

// TrapezoidRule is for solving the numerical integration using the trapezoid rule
func TrapezoidRule(a float64, b float64, N int, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / float64(N)
	x := a

	for i := 0; i < N; i++ {
		omega += f(x+h) + f(x)
		x += h
	}

	return h / 2 * omega
}

// SimpsonRule for solving numerical integration
func SimpsonRule(a float64, b float64, N int, f func(float64) float64) float64 {
	var omega float64
	h := (b - a) / float64(N)
	x := a

	for i := 0; i < N; i++ {
		omega += f(x) + 4*f(x+h) + f(x+2*h)
		x += h
	}
	return h / 6 * omega
}
