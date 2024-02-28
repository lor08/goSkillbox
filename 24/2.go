package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func parseTest(sentences []string, chars []rune) (positions [][]int) {
	sLen := len(sentences)

	maxNumberOfLetters := 0
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		lenLastWord := utf8.RuneCountInString(words[len(words)-1])
		if lenLastWord > maxNumberOfLetters {
			maxNumberOfLetters = lenLastWord
		}
	}
	cLen := maxNumberOfLetters

	positions = make([][]int, sLen)
	for i := range positions {
		positions[i] = make([]int, cLen)
	}

	for i, sentence := range sentences {
		words := strings.Split(sentence, " ")
		lastWord := words[len(words)-1]
		rLastWord := []rune(lastWord)

		for j, char := range chars {
			for letterIndex, letterInTheWord := range rLastWord {
				if letterInTheWord == char {
					positions[i][j] = letterIndex
				}
			}
		}
	}

	return
}

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет мама"}
	chars := []rune{'r', 'b', 'L', 'd', 'П', 'М', 'м'}

	positions := parseTest(sentences, chars)

	for i, sentence := range sentences {
		fmt.Println("Для предложения", sentence, "в последнем слове")
		words := strings.Split(sentence, " ")
		lastWord := words[len(words)-1]
		for j, char := range chars {
			if strings.Contains(lastWord, string(char)) {
				fmt.Printf("'%v' position %v\n", string(char), positions[i][j])
			}
		}
		fmt.Println()
	}
}
