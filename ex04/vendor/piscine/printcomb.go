package piscine

import (
	"ft"
)

func PrintComb() {
	for alp := '0'; alp <= '7'; alp++ {
		for alp2 := alp + 1; alp2 <= '8'; alp2 += 1 {
			for alp3 := alp2 + 1; alp3 <= '9'; alp3 += 1{
				ft.PrintRune(alp)
				ft.PrintRune(alp2)
				ft.PrintRune(alp3)
				if alp != '7' || alp2 != '8' || alp3 != '9' {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
	ft.PrintRune('\n')
}
