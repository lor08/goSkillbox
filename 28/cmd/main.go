package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"refactor28/pkg/check"
	"refactor28/pkg/storage"
	"refactor28/pkg/student"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	students := storage.New()
	i := 0
	for scanner.Scan() {
		if scanner.Err() == io.EOF {
			break
		}

		i++
		name, age, grade := check.InputParsing(scanner.Text())
		students.Put(student.New(name, age, grade))
	}

	fmt.Println()
	for _, student := range students {
		fmt.Println(student)
	}
}
