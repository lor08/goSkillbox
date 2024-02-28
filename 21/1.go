package main

import (
	"fmt"
	"math"
)

func calculation(x int16, y uint16, z float32) float32 {
	c := float64(y)
	return 2*float32(x) + float32(math.Pow(c, 2)) - 3/z
}

func main() {
	var (
		a int16   = 2
		b uint16  = 8
		c float32 = 12.345
	)

	s := calculation(a, b, c)

	fmt.Println(s)
}
