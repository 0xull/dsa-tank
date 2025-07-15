package main

import (
	"fmt"
	"math"
)

type term struct {
	coefficient float64
	exponent uint
	next *term
}

type polynomialExp struct {
	head *term
}

func (p *polynomialExp) addTerm(coeff float64, exp uint) {
	if coeff == 0.0 {
		return
	}
	
	if p.head == nil || exp > p.head.exponent {
		node := &term{coefficient: coeff, exponent: exp, next: p.head}
		p.head = node
		return
	} else if exp == p.head.exponent {
		sum := coeff + p.head.coefficient
		if sum != 0.0 {
			p.head.coefficient = sum
		} else {
			p.head = p.head.next
		}
		return
	}
	
	prev, curr := p.head, p.head.next
	for curr != nil && curr.exponent > exp {
		prev = curr
		curr = curr.next
	}
	
	if curr != nil && curr.exponent == exp {
		sum := curr.coefficient + coeff
		if sum != 0.0 {
			curr.coefficient = sum
		} else {
			prev.next = curr.next
		}
		
	} else {
		node := &term{coefficient: coeff, exponent: exp, next: curr}
		prev.next = node
	}
}

func (p *polynomialExp) display() {
	if p.head == nil {
		fmt.Println("0")
		return
	}
		
	fmt.Print("P(x) = ")
	firstTerm := true
	
	for curr := p.head; curr != nil; curr = curr.next {
		var term string
		exp := curr.exponent
		coeff := curr.coefficient
		
		if !firstTerm && coeff > 0 {
			term = " + "
		} else if coeff < 0 {
			if firstTerm {
				term = "-"
			} else {
				term = " - "
			}
			coeff = -coeff
		}
		
		if coeff != 1.0 || exp == 0 {
			if math.Mod(coeff, 1.0) == 0 {
				term = fmt.Sprintf("%s%.0f", term, coeff)
			} else {
				term = fmt.Sprintf("%s%.2f", term, coeff)
			}
		}
		
		if exp > 0 {
			term += "x"
			if exp > 1 {
				term += fmt.Sprintf("^%d", exp)
			} 
		}

		fmt.Print(term)
		firstTerm = false
	}
	
	fmt.Println("")
}

func (po *polynomialExp) add(p *polynomialExp) *polynomialExp {
	if po.head == nil {  // you said return a cop of p or po to avoid aliasing right, how do i return a copy?
		return p
	} else if p.head == nil {
		return po
	}
	
	newPoly := new(polynomialExp)
	tail := newPoly.head
	
	iterPo := po.head
	iterP := p.head
	
	for iterPo != nil && iterP != nil {
		var node *term
		if iterPo.exponent > iterP.exponent {
			node = &term{coefficient: iterPo.coefficient, exponent: iterPo.exponent}
			iterPo = iterPo.next
		} else if iterPo.exponent == iterP.exponent {
			sum := iterPo.coefficient + iterP.coefficient
			if sum != 0.0 {
				node = &term{coefficient: sum, exponent: iterPo.exponent}
			}
			iterP = iterP.next
			iterPo = iterPo.next
		} else {
			node = &term{coefficient: iterP.coefficient, exponent: iterP.exponent}
			iterP = iterP.next
		}
		
		if node != nil {
			if newPoly.head != nil {
				tail.next = node
				tail = tail.next
			} else {
				newPoly.head = node
				tail = newPoly.head
			}
		}
	}
	
	var remaining *term
	if iterPo != nil {
		remaining = iterPo
	} else {
		remaining = iterP
	}
	
	for remaining != nil {
		node := &term{coefficient: remaining.coefficient, exponent: remaining.exponent}
		if newPoly.head != nil {
			tail.next = node
			tail = tail.next
		} else {
			newPoly.head = node
			tail = newPoly.head
		}
		remaining = remaining.next
	}
	
	return newPoly
}

func main() {
    poly1 := &polynomialExp{}
    poly1.addTerm(5, 4)  // 5x^4
    poly1.addTerm(-3, 2) // -3x^2
    poly1.addTerm(1, 1)  // x
    poly1.addTerm(7, 0)  // 7
    fmt.Print("Polynomial 1: ")
    poly1.display() // Expected: 5x^4 - 3x^2 + x + 7

    poly2 := &polynomialExp{}
    poly2.addTerm(3, 3)  // 3x^3
    poly2.addTerm(3, 2)  // 3x^2 (will combine with -3x^2 from poly1 during addition)
    poly2.addTerm(-5, 0) // -5
    fmt.Print("Polynomial 2: ")
    poly2.display() // Expected: 3x^3 + 3x^2 - 5

    sumPoly := poly1.add(poly2)
    fmt.Print("Sum: ")
    sumPoly.display()
    // Expected: 5x^4 + 3x^3 + x + 2 (because -3x^2 + 3x^2 = 0x^2, so it's omitted)

    // Test adding a term that results in zero coefficient
    fmt.Println("\nTest zeroing out a term:")
    poly3 := &polynomialExp{}
    poly3.addTerm(2, 2)
    poly3.addTerm(5, 0)
    fmt.Print("Poly3 initial: ")
    poly3.display() // 2x^2 + 5
    poly3.addTerm(-2, 2) // Add -2x^2
    fmt.Print("Poly3 after adding -2x^2: ")
    poly3.display() // 5 (the x^2 term should be gone)
    
    poly4 := &polynomialExp{} // Empty polynomial
    poly4.addTerm(0,5) // Adding zero coefficient term
    fmt.Print("Poly4 (should be 0): ")
    poly4.display() // 0
}