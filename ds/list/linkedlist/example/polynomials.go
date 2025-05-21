package main


type term struct {
	coefficient float64
	exponent uint
	next *term
}

type polynomialExp struct {
	head *term
}

func (p *polynomialExp) AddTerm(coeff float64, exp uint) {
	if exp < 0 {
		return
	}
	
	if p.head == nil || exp > p.head.exponent {
		node := &term{coefficient: coeff, exponent: exp, next: p.head}
		p.head = node
	} else {
		prev, curr := (*term)(nil), p.head
		for curr != nil && curr.exponent > exp {
			prev = curr
			curr = curr.next
		}
		if curr.exponent == exp {
			curr.coefficient += coeff
		} else {
			node := &term{coefficient: coeff, exponent: exp, next: curr}
			prev.next = node
		}
	}
}

func main() {
	
}