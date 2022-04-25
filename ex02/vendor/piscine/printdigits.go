package piscine

import (
	"ft"
)

func PrintDigits() {
	for alp := '0'; alp <= '9'; alp++{
		ft.PrintRune(alp)
	}
	ft.PrintRune('\n')
}
