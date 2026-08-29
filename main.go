package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
input := "1, fg, 3,   4, 4.5, do   , -2, -7 , ha-ha, 17"

	// 1. Разбиваем строку на части по запятой
	parts := strings.Split(input, ",")
	fmt.Println(parts)

	fmt.Println("___ Калькулятор (среднего, суммы, медиального) ___")
	operation := getOperation()
	numbersSlice := getNumbersSlice()
	finishNumber := getOperationWithSlice(operation, numbersSlice)
	
	fmt.Printf("Ваше искомое число: %v", finishNumber)
}

func getOperation() string {
	var operation string
	for {
		fmt.Print("Введите операцию. 'AVG', 'SUM' или 'MID': ")
		_, err := fmt.Scan(&operation)
		if err != nil {
			fmt.Println("Ошибка ввода.")
			continue
		}
		if operation == "AVG" || operation == "SUM" || operation == "MID" {
			return operation
		}
		fmt.Println("Ошибка ввода. Попробуйте ещё раз.")
	}
}

func getNumbersSlice() []float64 {
	fmt.Println("Введите числа через запятую: ")
	// 1. Принимает любую строку со всеми пробелами, в отличии от Scan(перед пробелом останавлтвант считывание)
	reader := bufio.NewReader(os.Stdin)
	inputUser, _ := reader.ReadString('\n')

	parts := strings.Split(inputUser, ",")
	numbers := []float64 {}
	for _, part := range parts {
		// 2. Очищаем от пробелов в начале и конце (" fg" -> "fg", " 4.5" -> "4.5")
		trimmed := strings.TrimSpace(part)
		
		// 3. Пробуем преобразовать строку в число типа float64
		if num, err := strconv.ParseFloat(trimmed, 64); err == nil {
			// Если ошибки нет, добавляем число в наш слайс
			numbers = append(numbers, num)
		}
	}
	slices.Sort(numbers)
	return numbers
}

func getOperationWithSlice(operation string, slice []float64) float64 {
	count := 0.0
	n := len(slice)
	switch operation {
	case "AVG": 
		for _, value := range slice {
			count += value
		}
		count = count / float64(n)
	case "SUM":
		for _, value := range slice {
			count += value
		}
	case "MID":
		if n % 2 != 0 {
		// Количество нечётное (например, 7 элементов): берем центральный (индекс 3)
		count = slice[n / 2]
		} else {
			// Количество чётное: берем два центральных и находим среднее
			middle1 := slice[(n/2)-1]
			middle2 := slice[n/2]
			count = (middle1 + middle2) / 2.0
		}
	}
	return count
}