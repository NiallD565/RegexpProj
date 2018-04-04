package util

import (
	"fmt"
)

func PrintS(s string) {
	fmt.Println(s)
}

// Input take input from the user and assigns it to s which is needed by the nfa
func Input(s string) string {
	var Input string
	PrintS(s)
	fmt.Scan(&Input)
	return Input

}
