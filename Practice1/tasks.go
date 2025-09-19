package main

import (
	"fmt"
	"time"
)

func sumAndDiff(a, b float64) (float64, float64) { //функция для 5 задания
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	//1 task
	currentTime := time.Now()
	formattedTime := currentTime.Format("02-01-2006 15:04:05")
	fmt.Println("1 Задание.")
	fmt.Println("Текущее время и дата:", formattedTime)
	fmt.Println("==================================================")

	//2 task
	var name string = "Artem"
	var age int = 21
	var isStudent bool = true
	var avgMark float64 = 4.98

	fmt.Println("2 Задание.")
	fmt.Println("Имя:", name)
	fmt.Println("Возврат:", age)
	fmt.Println("Студент:", isStudent)
	fmt.Println("Средний балл:", avgMark)
	fmt.Println("==================================================")

	//3 task
	fmt.Println("3 Задание.")
	name = "Anton"
	surname := "Nikolaev"
	age = 21
	isStudent = true
	avgMark = 4.88
	fmt.Println("Имя:", name)
	fmt.Println("Фамилия:", surname)
	fmt.Println("Возврат:", age)
	fmt.Println("Студент:", isStudent)
	fmt.Println("Средний балл:", avgMark)
	fmt.Println("==================================================")

	//4 task
	fmt.Println("4 Задание.")
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	umnozh := a * b
	quotient := a / b
	remainder := a % b
	fmt.Println("a =", a, "b =", b)
	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)
	fmt.Println("Произведение:", umnozh)
	fmt.Println("Деление:", quotient)
	fmt.Println("Остаток от деления:", remainder)
	fmt.Println("==================================================")

	//5 task
	fmt.Println("5 Задание.")
	num1 := 12.5
	num2 := 7.3
	sum1, diff1 := sumAndDiff(num1, num2)
	fmt.Println("Числа:", num1, "и", num2)
	fmt.Println("Сумма:", sum1)
	fmt.Println("Разность:", diff1)
	fmt.Println("==================================================")

	//6 task
	var num3, num4, num5 float64
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num3)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num4)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&num5)
	avg := (num3 + num4 + num5) / 3
	fmt.Println("Среднее значение:", avg)
}
