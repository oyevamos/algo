package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("INPUT.TXT")
	// min := 111111
	// max := -111111
	if err != nil {
		fmt.Println("Ошибка при чтении файла")
		return
	}
	defer file.Close()

	min := math.MaxInt64
	max := math.MinInt64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		// number, err := strconv.Atoi(line)
		for _, numStr := range numbers {
			number, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Ошибка при преобразовании строки в число")
				fmt.Println("Проблемная строка:", line)
				continue
			}
			min = int(math.Min(float64(min), float64(number)))
			max = int(math.Max(float64(max), float64(number)))
		}
	}
	outputFile, err := os.Create("OUTPUT.TXT")
	if err != nil {
		fmt.Println("Ошибка при создании файла для записи:", err)
		return
	}
	defer outputFile.Close()
	_, err = fmt.Fprintln(outputFile, max-min)
	if err != nil {
		fmt.Println("Ошибка при записи данных в файл вывода:", err)
		return
	}

}
