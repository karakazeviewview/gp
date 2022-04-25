package piscine

import (
	"ft"
)

func PrintReverseAlphabet() {
	for alp := 'z'; alp >= 'a'; alp--{
		ft.PrintRune(alp)
	}
	ft.PrintRune('\n')
}
