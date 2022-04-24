package piscine

import (
	"ft"
)

func Printreversealphabet() {
	for alp := 'z'; alp >= 'a'; alp--{
		ft.PrintRune(alp)
	}
	ft.PrintRune('\n')
}
