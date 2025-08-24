package test

import "github.com/01-edu/z01"

func Print(str ...string) {
	for _, v := range str {
		for _, s := range v {
			z01.PrintRune(s)
		}
		z01.PrintRune(' ')
	}
}

func Println(str ...string) {
	for _, v := range str {
		for _, s := range v {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
