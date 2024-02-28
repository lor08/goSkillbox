package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func findSubstrWithContains(str, substr *string) (bool, error) {
	if str == nil || substr == nil {
		return false, fmt.Errorf("nil pointer")
	}
	return strings.Contains(*str, *substr), nil
}

func findSubstrWithRune(str, substr *string) (bool, error) {
	if str == nil || substr == nil {
		return false, fmt.Errorf("nil pointer")
	}

	rStr := []rune(*str)
	rSubStr := []rune(*substr)

str:
	for i, charStr := range rStr {
		if charStr == rSubStr[0] {
			for j, charSubStr := range rSubStr[1:] {
				if rStr[i+j+1] != charSubStr {
					continue str
				}
			}
			return true, nil
		}
	}

	return false, nil
}

func displayResult(isFind bool, err error) {
	if err != nil {
		log.Fatalln("Ошибка приёма данных\nerr:", err)
	} else {
		if isFind {
			fmt.Println("Есть вхождение подстроки.")
		} else {
			fmt.Println("Вхождение не найдено")
		}
	}
}

func main() {
	var str, substr string
	flag.StringVar(&str, "str", "", "usage str")
	flag.StringVar(&substr, "substr", "", "usage substr")
	flag.Parse()

	fmt.Println("Поиск подстроки с помощью strings.Contains")
	isFind, err := findSubstrWithContains(&str, &substr)
	displayResult(isFind, err)

	fmt.Println("\nПоиск подстроки с помощью самописного алгоритма")
	isFind, err = findSubstrWithRune(&str, &substr)
	displayResult(isFind, err)
}
