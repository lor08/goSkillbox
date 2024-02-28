package main

import "fmt"

func calc(t string, a, b int, A func(int, int) int) {
	defer fmt.Println("результат:", A(a, b))
	fmt.Println()
	fmt.Printf("%v (%v, %v)\n", t, a, b)
}

func main() {
	x := 5
	y := 2
	calc("Сложение", x, y, func(x, y int) int { return x + y })
	calc("Вычитание", x, y, func(x, y int) int { return x - y })
	calc("Умножение", x, y, func(x, y int) int { return x * y })
}
