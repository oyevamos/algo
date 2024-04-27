package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("INPUT.TXT")
	if err != nil {
		fmt.Println("Ошибка при открытии файла", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()

	oldestAge := 0
	oldestIndex := -1
	index := 1

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		age, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Ошибка при чтении возраста", err)
			continue
		}
		gender, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Ошибка при чтении пола:", err)
			continue
		}
		if gender == 1 && age > oldestAge {
			oldestAge = age
			oldestIndex = index
		}
		index++
	}
	outputFile, err := os.Create("OUTPUT.TXT")
	if err != nil {
		fmt.Println("Ошибка при создании файла для записи:", err)
		return
	}
	defer outputFile.Close()

	if oldestIndex == -1 {
		fmt.Fprintln(outputFile, -1)
	} else {
		fmt.Fprintln(outputFile, oldestIndex)
	}
}
