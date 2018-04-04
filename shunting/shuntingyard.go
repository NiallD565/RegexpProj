package shunting

func intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	// * 0 or more
	// + 1 or more
	// . concatenate
	// | or

	postfix := []rune{}
	s := []rune{}

	// ( added to stack
	// ) add each character stack to the postfix while the opening tag not found get rid of the bracket
	for _, r := range infix { // convert string array to runes
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				postfix = append(postfix, s[len(s)-1])
				s = s[:len(s)-1] // get everything but the last element in the list
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				postfix = append(postfix, s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = append(s, r)
		default:
			postfix = append(postfix, r)
		}
	}
	for len(s) > 0 { // if any characters are left append them to the string and clear the stack
		postfix, s = append(postfix, s[len(s)-1]), s[:len(s)-1] // get everything but the last element in the list
	}

	return string(postfix)
}
