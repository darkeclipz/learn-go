package foohandler

import (
	"fmt"
	"log"
	"time"
)

type variable struct {
	id    int
	value int
}

type propagation struct {
	variable, value int
}

type notEqualConstraint struct {
	left, right int
}

type mapColoringCsp struct {
	variables   []variable
	constraints []notEqualConstraint
}

const (
	WA           = 0
	NT           = 1
	SA           = 2
	Q            = 3
	NSW          = 4
	V            = 5
	T            = 6
	Unassigned   = -1
	NumVariables = 7
	NumColors    = 3
)

var (
	csp = mapColoringCsp{
		variables: []variable{
			{WA, Unassigned},
			{NT, Unassigned},
			{SA, Unassigned},
			{Q, Unassigned},
			{NSW, Unassigned},
			{V, Unassigned},
			{T, Unassigned},
		},
		constraints: []notEqualConstraint{
			{WA, NT}, {WA, SA},
			{NT, SA}, {NT, Q},
			{SA, Q}, {SA, NSW}, {SA, V},
			{Q, NSW},
			{NSW, V},
		},
	}
	colorTaken   [NumVariables][3]bool
	propagations [NumVariables][]propagation
)

func main() {
	defer duration(track("solve"))
	solve(0)
	print()
}

func solve(current int) bool {
	if current >= len(csp.variables) {
		return true
	}
	for color := range NumColors {
		if colorTaken[current][color] {
			//fmt.Printf("skipping color %d.\n", color)
			continue
		}
		//fmt.Printf("assign %v=%v\n", current, color)
		csp.variables[current].value = color
		propagate(current, color)
		if valid() && solve(current+1) {
			return true
		}
		//fmt.Printf("unassign %v\n", current)
		csp.variables[current].value = Unassigned
		undoPropagation(current)
	}
	return false
}

func valid() bool {
	for _, c := range csp.constraints {
		if c.inScope() && !c.satisfied() {
			return false
		}
	}
	return true
}

func (c notEqualConstraint) satisfied() bool {
	return csp.variables[c.left].value != csp.variables[c.right].value
}

func (c notEqualConstraint) inScope() bool {
	if csp.variables[c.left].value == Unassigned {
		return false
	}
	if csp.variables[c.right].value == Unassigned {
		return false
	}
	return true
}

func propagate(current int, value int) {
	propagations[current] = make([]propagation, NumVariables)
	for _, cc := range csp.constraints {
		if cc.left == current && !colorTaken[cc.right][value] {
			colorTaken[cc.right][value] = true
			//fmt.Printf("propagate curr=%v, other=%v, color=%v\n", current, cc.right, value)
			propagations[current] = append(propagations[current], propagation{cc.right, value})
		}
	}
}

func undoPropagation(current int) {
	for _, prop := range propagations[current] {
		colorTaken[prop.variable][prop.value] = false
	}
}

func print() {
	for i, v := range csp.variables {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
