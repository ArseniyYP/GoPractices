package main

import "fmt"

func checkNumber(n int) string { //for 2nd task
	if n > 0 {
		return "Positive"
	} else if n < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func getLength(s string) int { //for 4 task
	return len(s)
}

type Rectangle struct { //for 5 task
	width  float64
	height float64
}

func (r Rectangle) Area() float64 { //for 5 task
	return r.width * r.height
}

func averange(a, b int) float64 { //for 6 task
	return float64(a+b) / 2
}

func main() {
	//1 task
	fmt.Println("1 Задание.")
	var num1 int
	fmt.Print("Введите число: ")
	fmt.Scan(&num1)
	if num1%2 == 0 {
		fmt.Println("Число", num1, "- четное")
	} else {
		fmt.Println("ЧИсло", num1, "- нечетное")
	}
	fmt.Println("==================================================")

	//2 task
	fmt.Println("2 Задание.")
	var num2 int
	fmt.Print("Введите число: ")
	fmt.Scan(&num2)
	result2 := checkNumber(num2)
	fmt.Println("Результат:", result2)
	fmt.Println("==================================================")

	//3 task
	fmt.Println("3 Задание.")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println("==================================================")

	//4 task
	fmt.Println("4 Задание.")
	var text4 string
	fmt.Print("Введите строку: ")
	fmt.Scan(&text4)
	length4 := getLength(text4)
	fmt.Println("Длина строки:", length4)
	fmt.Println("==================================================")

	//5 task
	fmt.Println("5 Задание.")
	rect5 := Rectangle{width: 5, height: 3}
	fmt.Println("Ширина:", rect5.width)
	fmt.Println("Высота:", rect5.height)
	fmt.Println("Площадь:", rect5.Area())
	fmt.Println("==================================================")

	//6 task
	fmt.Println("6 Задание.")
	var x, y int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)
	result6 := averange(x, y)
	fmt.Println("Среднее значение:", result6)
}
