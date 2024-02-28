package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	var numbers [size]int
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(10 * size)
	}

	fmt.Println("Заданный массив")
	fmt.Printf("%+v\n", numbers)

	fmt.Println("Введите число, после которого подсчитаем кол-во чисел в массиве")
	n := 0
	fmt.Scan(&n)

	index := 0
	for i := 0; i < size; i++ {
		if numbers[i] == n {
			index = i
		}
	}
	count := len(numbers) - index - 1

	fmt.Println("Количество чисел в массиве после указанного:", count)
}
