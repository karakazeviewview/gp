package piscine

import (
	"ft"
)

func PrintAlphabet() {
	for alp := 'a'; alp <= 'z'; alp++{
		ft.PrintRune(alp)
	}
	ft.PrintRune('\n')
}
