package main

import (
	"fmt"
)

// test github desktop
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

func poregtonfa(postfix string) *nfa {
	nfastack := []*nfa{}

	for _, r := range postfix {
		switch r {
		case '.':
			// pointers to nfa fragments
			frag2 := nfastack[len(nfastack)-1] // index of the last item on the stack
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1] // index of the last item on the stack
			nfastack = nfastack[:len(nfastack)-1]

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
			frag := nfastack[len(nfastack)-1] // index of the last item on the stack
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}
	return nfastack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
