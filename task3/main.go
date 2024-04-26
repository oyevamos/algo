package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("INPUT.TXT")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	defer file.Close()

	cnt := 0
	max := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if char == '0' {
				cnt++
				if cnt > max {
					max = cnt
				}
			} else {
				cnt = 0
			}
		}
	}

	outputFile, err := os.Create("OUTPUT.TXT")
	if err != nil {
		fmt.Println("Ошибка при создании файла для записи:", err)
		return
	}
	defer outputFile.Close()

	if _, err = fmt.Fprintln(outputFile, max); err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}
