package main

import (
	"fmt"
	"math"
)

type Person struct { //for 1st task
	name string
	age  int
}

func (p Person) Info() { //for 1st task
	fmt.Printf("Имя: %s, Возраст: %d\n", p.name, p.age)
}

func (p *Person) Birthday() { //for 2nd task
	p.age++
}

type Circle struct { //for 3rd task
	radius float64
}

func (c Circle) Area() float64 { //for 3rd task
	return math.Pi * c.radius * c.radius
}

type Shape interface { //for 4th task
	Area() float64
}

type Rectangle struct { //for 4th task
	width, height float64
}

func (r Rectangle) Area() float64 { //for 4th task
	return r.width * r.height
}

func PrintAreas(shapes []Shape) { //for 5th task
	for i, shape := range shapes {
		fmt.Printf("Фигура %d: площадь = %.2f\n", i+1, shape.Area())
	}
}

type Book struct { //for 6th task
	title  string
	author string
	pages  int
}

func (b Book) String() string { //for 6th task
	return fmt.Sprintf("Книга: \"%s\", Автор: %s, Страниц: %d", b.title, b.author, b.pages)
}

func main() {
	//1 task
	fmt.Println("1 Задание.")
	person1 := Person{name: "Алексей", age: 25}
	person2 := Person{name: "Мария", age: 30}
	person1.Info()
	person2.Info()
	fmt.Println("==================================================")

	//2 task
	fmt.Println("2 Задание.")
	fmt.Print("Информация до дня рождения: ")
	person1.Info()
	fmt.Print("Празднуем день рождение")
	person1.Birthday()
	fmt.Print("Информация после д/р: ")
	person1.Info()
	fmt.Println("==================================================")

	//3 task
	fmt.Println("3 Задание.")
	circle := Circle{radius: 5}
	fmt.Printf("Радиус круга: %.2f\n", circle.radius)
	fmt.Printf("Площадь круга: %.2f\n", circle.Area())
	fmt.Println("==================================================")

	//4 task
	fmt.Println("4 Задание.")
	r := Rectangle{width: 4, height: 5}
	c := Circle{radius: 3}
	shapes := []Shape{r, c}
	for _, shape := range shapes {
		fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
	}
	fmt.Println("==================================================")

	//5 task
	fmt.Println("5 Задание.")
	PrintAreas(shapes)
	fmt.Println("==================================================")

	//6 task
	fmt.Println("6 Задание.")
	book := Book{
		title:  "Война и мир",
		author: "Лев Толстой",
		pages:  1225,
	}
	fmt.Println(book)
	fmt.Println("==================================================")
}
