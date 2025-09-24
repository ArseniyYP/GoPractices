package main

import (
	"fmt"
	"practice3/mathutils"
	"practice3/stringutils"
)

func main() {
	//1-2 task
	fmt.Println("1-2 Задания.")
	var num1 int
	fmt.Print("Введите число: ")
	fmt.Scan(&num1)
	fmt.Printf("Факториал %d = %d\n", num1, mathutils.Factorial(num1))
	fmt.Println("==================================================")

	//3 task
	fmt.Println("3 Задание.")
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scan(&input)
	result3 := stringutils.Reverse(input)
	fmt.Println("Перевернутая строка:", result3)
	fmt.Println("==================================================")

	//4 task
	fmt.Println("4 Задание.")
	var arr4 [5]int
	arr4[0] = 10
	arr4[1] = 20
	arr4[2] = 30
	arr4[3] = 40
	arr4[4] = 50
	fmt.Println("Массив:", arr4)
	fmt.Println("Элементы массива:")
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr4[i])
	}
	fmt.Println("==================================================")

	//5 task
	fmt.Println("5 Задание.")
	arr5 := [5]int{10, 20, 30, 40, 50}
	slice := arr5[1:4]
	fmt.Println("Начальный срез:", slice)
	slice = append(slice, 60)
	fmt.Println("После добавления:", slice)
	index5 := 1
	slice = append(slice[:index5], slice[index5+1:]...)
	fmt.Println("После удаления:", slice)
	fmt.Println("==================================================")

	//6 task
	fmt.Println("6 Задание.")
	words := []string{"Go", "Programming", "Language", "Slice", "Example"}
	longest := words[0]
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	fmt.Println("Срез строк:", words)
	fmt.Println("Самая длинная строка:", longest)
	fmt.Println("==================================================")
}
