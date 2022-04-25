package piscine

import (
	"ft"
)

func PrintComb2() {
	for alp := '0'; alp <= '9'; alp++ {
		for alp2 := '0'; alp2 <= '9'; alp2 += 1 {
			for alp3 := '0'; alp3 <= '9'; alp3 += 1 {
				for alp4 := '0'; alp4 <= '9'; alp4 += 1 {
					if alp <= alp3 && (alp != alp3 || alp2 < alp4){
						ft.PrintRune(alp)
						ft.PrintRune(alp2)
						ft.PrintRune(' ')
						ft.PrintRune(alp3)
						ft.PrintRune(alp4)
						if alp != '9' ||  alp2 != '8' || alp3 != '9' || alp4 != '9' {
							ft.PrintRune(',')
							ft.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	ft.PrintRune('\n')
}
