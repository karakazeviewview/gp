package piscine

import (
	"ft"
)

func printalphabet() {

	i := 0
	for i < 25 {
		alp := 'a' + i;
		ft.PrintRune(alp)
		i++
	}
	return 0
}
