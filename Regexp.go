package main

import (
	"fmt"
)

// pointers to other states
type state struct {
	symbol rune // rune gives the integer value of a character
	edge1  *state
	edge2  *state
}

// keeps track of the initial state and the accept state of the fragment
type nfa struct {
	initial *state
	accept  *state
}

// uses a stack
func poregtonfa(postfix string) *nfa {
	nfastack := []*nfa{}
	// loop through the postfix
	// uses r to represent the given character from the string
	for _, r := range postfix {
		switch r {
		case '.': //concatination
			// Array to pointers to nfa fragments
			// Pops two items off the nfa stack
			frag2 := nfastack[len(nfastack)-1]    // index of the last item on the stack
			nfastack = nfastack[:len(nfastack)-1] // : gives everything up not including the lasr item
			frag1 := nfastack[len(nfastack)-1]    // index of the last item on the stack
			nfastack = nfastack[:len(nfastack)-1]

			// edge1 points to the initial state of frag2
			frag1.accept.edge1 = frag2.initial

			// & gives the address of the instance
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			// pointers to nfa fragments
			frag2 := nfastack[len(nfastack)-1] // index of the last item on the stack
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1] // index of the last item on the stack
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// & gives the address of the instance
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*':
			// only need to pop 1 frag off the stack because * only works on 1 fragments
			frag := nfastack[len(nfastack)-1] // index of the last item on the stack
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial // joins the accept state to the initial state
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		default:
			// no need to pop anthing from the stack
			accept := state{}
			initial := state{symbol: r, edge1: &accept} // set the symbol to r or else it will have its default value of 0

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}
	if len(nfastack) != 1 {
		fmt.Println("Nani:", len(nfastack), nfastack)
	}

	return nfastack[0] // nfa return value
}

// get the current array
// add s to it and check if it is a possible initial state
func addstate(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	// if the state has a E arrow to it, it is assigned the value of 0
	if s != a && s.symbol == 0 {
		l = addstate(l, s.edge1, a)
		// check for a second edge
		if s.edge2 != nil {
			l = addstate(l, s.edge2, a)
		}
	}
	return l
}

func postmatch(po string, s string) bool {
	ismatch := false
	postnfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	// list containing all the possible initial states
	current = addstate(current[:], postnfa.initial, postnfa.accept)

	for _, r := range s {
		for _, c := range current {
			// loop through the current states naming them c
			// check the state if the name is set to r
			if c.symbol == r {
				// r has a single arrow going from it to another another state
				next = addstate(next[:], c.edge1, postnfa.accept)
			}
		}
		current, next = next, []*state{}
	}
	for _, c := range current {
		if c == postnfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func main() {
	nfa := poregtonfa("ab.c*|") // test case representing a regular expression
	fmt.Println(nfa)            // print out the return value (memory address)

	fmt.Println(postmatch("ab.c*|", "cccc"))
}
