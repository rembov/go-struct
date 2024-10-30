package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	switch p.Age {
	case 1, 21, 31, 41, 51, 61, 71, 81, 91:
		fmt.Printf("Привет, меня зовут %s и мне %d год\n", p.Name, p.Age)
	case 2, 3, 4, 22, 23, 24, 32, 33, 34, 42, 43, 44, 52, 53, 54, 62, 63, 64, 72, 73, 74, 82, 83, 84, 92, 93, 94:
		fmt.Printf("Привет, меня зовут %s и мне %d года\n", p.Name, p.Age)
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 25, 26, 27, 28, 29, 30, 35, 36, 37, 38, 39, 40, 45, 46, 47, 48, 49, 50, 55, 56, 57, 58, 59, 60, 65, 66, 67, 68, 69, 70, 75, 76, 77, 78, 79, 80, 95, 96, 97, 98, 99, 100:
		fmt.Printf("Привет, меня зовут %s и мне %d лет\n", p.Name, p.Age)
	default:
		fmt.Println("Столько не живут)))")
	}

}

func (p *Person) HaveBirthday() {
	p.Age++
	switch p.Age {
	case 1, 21, 31, 41, 51, 61, 71, 81, 91:
		fmt.Printf("Теперь мне %d год\n", p.Age)
	case 2, 3, 4, 22, 23, 24, 32, 33, 34, 42, 43, 44, 52, 53, 54, 62, 63, 64, 72, 73, 74, 82, 83, 84, 92, 93, 94:
		fmt.Printf("Теперь мне %d года\n", p.Age)
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 25, 26, 27, 28, 29, 30, 35, 36, 37, 38, 39, 40, 45, 46, 47, 48, 49, 50, 55, 56, 57, 58, 59, 60, 65, 66, 67, 68, 69, 70, 75, 76, 77, 78, 79, 80, 95, 96, 97, 98, 99, 100:
		fmt.Printf("Теперь мне %d лет\n", p.Age)
	default:
		fmt.Println("Столько не живут)))")

	}
}

func main() {
	v := 0
	name := ""
	fmt.Println("Введите возраст")
	fmt.Scan(&v)

	fmt.Println("Введите имя")
	fmt.Scan(&name)
	person1 := Person{Name: name, Age: v}

	person1.Greet()

	person1.HaveBirthday()
}
