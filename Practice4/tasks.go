package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func averangeAge(people map[string]int) float64 { //for 2nd task
	if len(people) == 0 {
		return 0
	}
	sum := 0
	for _, age := range people {
		sum += age
	}
	return float64(sum) / float64(len(people))
}

func main() {
	//1 task
	fmt.Println("1 Задание.")
	people := map[string]int{
		"Данил": 25,
		"Антон": 30,
		"Артем": 22,
	}
	people["Олег"] = 28
	fmt.Println("Список людей:")
	for name, age := range people {
		fmt.Printf("%s - %d лет\n", name, age)
	}
	fmt.Println("==================================================")

	//2 task
	fmt.Println("2 Задание.")
	avg := averangeAge(people)
	fmt.Println("Список людей:", people)
	fmt.Printf("Средний возраст: %.2f\n", avg)
	fmt.Println("==================================================")

	//3 task
	fmt.Println("3 Задание.")
	fmt.Println("До удаления:", people)
	nameToDelete := "Антон"
	delete(people, nameToDelete)
	fmt.Printf("После удаления %s: %v\n", nameToDelete, people)
	fmt.Println("==================================================")

	//4 task
	fmt.Println("4 Задание.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n') //считывание строки с пробелами
	input = strings.TrimSpace(input)    //убирает лишние пробелы и перевод строки
	upper := strings.ToUpper(input)     //перевод в верхний регистр
	fmt.Println("В верхнем регистре:", upper)
	fmt.Println("==================================================")

	//5 task
	fmt.Println("5 Задание.")
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через пробел: ")
	input1, _ := reader1.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	parts := strings.Split(input1, " ")
	sum := 0
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Ошибка: введено не число -", part)
			continue
		}
		sum += num
	}
	fmt.Println("Сумма чисел:", sum)
	fmt.Println("==================================================")

	//6 task
	fmt.Println("6 Задание.")
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Введите целые числа через пробел: ")
	input2, _ := reader2.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	parts1 := strings.Split(input2, " ")
	var numbers []int
	for _, part1 := range parts1 {
		num, err := strconv.Atoi(part1)
		if err != nil {
			fmt.Println("Ошибка: введено не число -", part1)
			continue
		}
		numbers = append(numbers, num)
	}
	fmt.Println("Исходный массив:", numbers)
	fmt.Print("Массив в обратной порядке: ")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()
	fmt.Println("==================================================")
}
