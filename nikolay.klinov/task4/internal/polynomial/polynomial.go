package polynomial

import "fmt"

type Polynomial struct {
	Coefficients []bool
	Degree       int
}

func New(degree int, bits uint64) *Polynomial {
	coeffs := make([]bool, degree+1)
	for i := 0; i <= degree; i++ {
		coeffs[i] = (bits>>uint(i))&1 == 1
	}
	return &Polynomial{
		Coefficients: coeffs,
		Degree:       degree,
	}
}

func (p *Polynomial) ToUint64() uint64 {
	var result uint64
	for i, coeff := range p.Coefficients {
		if coeff {
			result |= 1 << uint(i)
		}
	}
	return result
}

func (p *Polynomial) String() string {
	s := ""
	for i := p.Degree; i >= 0; i-- {
		if p.Coefficients[i] {
			if i > 1 {
				s += fmt.Sprintf("x^%d + ", i)
			} else if i == 1 {
				s += "x + "
			} else {
				s += "1 + "
			}
		}
	}
	if len(s) > 0 {
		s = s[:len(s)-3]
	}
	return s
}
