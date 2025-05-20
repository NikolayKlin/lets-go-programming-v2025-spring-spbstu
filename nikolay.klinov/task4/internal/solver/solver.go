package solver

import (
	"sync"
	"task4/internal/polynomial"
)

type Solver struct {
	mu    sync.Mutex
	wg    sync.WaitGroup
	stop  chan struct{}
	found chan *polynomial.Polynomial
}

func NewSolver() *Solver {
	return &Solver{
		stop:  make(chan struct{}),
		found: make(chan *polynomial.Polynomial, 1),
	}
}

func (s *Solver) FindIrreducibleAsync(degree int) *polynomial.Polynomial {
	start := uint64(1 << uint(degree))
	end := uint64(1 << uint(degree+1))

	for bits := start; bits < end; bits++ {
		select {
		case <-s.stop:
			return nil
		default:
			p := polynomial.New(degree, bits)
			if s.isIrreducible(p) {
				return p
			}
		}
	}
	return nil
}

func (s *Solver) FindIrreducibleSync(degree int) *polynomial.Polynomial {
	start := uint64(1 << uint(degree))
	end := uint64(1 << uint(degree+1))
	tasks := make(chan uint64, 100)

	for i := 0; i < 24; i++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			for {
				select {
				case <-s.stop:
					return
				case bits, ok := <-tasks:
					if !ok {
						return
					}
					p := polynomial.New(degree, bits)
					if s.isIrreducible(p) {
						s.mu.Lock()
						select {
						case s.found <- p:
							close(s.stop)
						default:
						}
						s.mu.Unlock()
					}
				}
			}
		}()
	}

	go func() {
		defer close(tasks)
		for bits := start; bits < end; bits++ {
			select {
			case <-s.stop:
				return
			case tasks <- bits:
			}
		}
	}()

	s.wg.Wait()

	select {
	case res := <-s.found:
		return res
	default:
		return nil
	}
}

func (s *Solver) isIrreducible(p *polynomial.Polynomial) bool {
	for d := 1; d <= p.Degree/2; d++ {
		select {
		case <-s.stop:
			return false
		default:
			if s.hasDivisor(p, d) {
				return false
			}
		}
	}
	return true
}

func (s *Solver) hasDivisor(p *polynomial.Polynomial, degree int) bool {
	start := uint64(1 << uint(degree))
	end := uint64(1 << uint(degree+1))

	for bits := start; bits < end; bits++ {
		select {
		case <-s.stop:
			return false
		default:
			divisor := polynomial.New(degree, bits)
			if isDivisor(p, divisor) {
				return true
			}
		}
	}
	return false
}

func isDivisor(a, b *polynomial.Polynomial) bool {
	if b.Degree == 0 {
		return true
	}
	if a.Degree < b.Degree {
		return false
	}
	remainder := clone(a)
	for remainder.Degree >= b.Degree {
		shift := remainder.Degree - b.Degree
		shifted := shiftLeft(b, shift)
		remainder = add(remainder, shifted)
	}
	return remainder.Degree == -1
}

func clone(p *polynomial.Polynomial) *polynomial.Polynomial {
	coeffs := make([]bool, len(p.Coefficients))
	copy(coeffs, p.Coefficients)
	return &polynomial.Polynomial{
		Coefficients: coeffs,
		Degree:       p.Degree,
	}
}

func shiftLeft(p *polynomial.Polynomial, shift int) *polynomial.Polynomial {
	coeffs := make([]bool, p.Degree+shift+1)
	for i := 0; i <= p.Degree; i++ {
		coeffs[i+shift] = p.Coefficients[i]
	}
	return &polynomial.Polynomial{
		Coefficients: coeffs,
		Degree:       p.Degree + shift,
	}
}

func add(a, b *polynomial.Polynomial) *polynomial.Polynomial {
	maxDegree := a.Degree
	if b.Degree > maxDegree {
		maxDegree = b.Degree
	}
	coeffs := make([]bool, maxDegree+1)
	for i := 0; i <= a.Degree; i++ {
		coeffs[i] = coeffs[i] != a.Coefficients[i]
	}
	for i := 0; i <= b.Degree; i++ {
		coeffs[i] = coeffs[i] != b.Coefficients[i]
	}
	degree := -1
	for i := len(coeffs) - 1; i >= 0; i-- {
		if coeffs[i] {
			degree = i
			break
		}
	}
	return &polynomial.Polynomial{
		Coefficients: coeffs[:degree+1],
		Degree:       degree,
	}
}
