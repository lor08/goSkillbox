package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sizeN = 12

func find(a [sizeN]int, value int) (index int) {
	index = -1
	min := 0
	max := sizeN - 1
	for max >= min {
		middle := (max + min) / 2
		if a[middle] == value {
			index = middle
			break
		} else if a[middle] > value {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [sizeN]int

	for i := 0; i < sizeN; i++ {
		a[i] = 10*i + rand.Intn(10)
	}

	fmt.Println("Исходный массив")
	fmt.Println(a)

	value := a[rand.Intn(sizeN)]
	index := find(a, value)
	fmt.Println("Индекс:", index)
}
