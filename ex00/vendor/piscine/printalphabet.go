package piscine

import (
	"ft"
)

func Printalphabet() {
	for alp := 'a'; alp <= 'z'; alp++{
		ft.PrintRune(alp)
	}
	ft.PrintRune('\n')
}
