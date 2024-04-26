package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("INPUT.TXT")
	if err != nil {
		fmt.Println("Ошибка при чтении файла")
		return
	}
	defer file.Close()

	cnt := 0
	max := -11111

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			// fmt.Println(string(line[i]))
			if string(line[i]) == "0" {
				cnt++
				max = int(math.Max(float64(max), float64(cnt)))
			} else {
				cnt = 0
			}
		}
	}
	// fmt.Println(max)
	outputFile, err := os.Create("OUTPUT.TXT")
	if err != nil {
		fmt.Println("Ошибка при создании файла для записи ", err)
		return
	}
	defer outputFile.Close()
	// buba := strconv.Itoa(max)
	_, err = fmt.Fprintln(outputFile, max)
}
