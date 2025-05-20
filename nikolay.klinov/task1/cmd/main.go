package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ввод первого числа
	fmt.Print("Введите первое число: ")
	num1Str, _ := reader.ReadString('\n')
	num1Str = strings.TrimSpace(num1Str)
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil { // Обработка ввода
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	// Ввод операции
	fmt.Print("Выберите операцию (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)
	if op != "+" && op != "-" && op != "*" && op != "/" { // Обработка ввода
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		return
	}

	// Ввод второго числа
	fmt.Print("Введите второе число: ")
	num2Str, _ := reader.ReadString('\n')
	num2Str = strings.TrimSpace(num2Str)
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil { // Обработка ввода
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	// Выполнение операции и вывод результата
	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 { // Обработка деления на 0
			fmt.Println("Ошибка: деление на ноль невозможно.")
			return
		}
		result = num1 / num2
	}

	fmt.Printf("Результат: %v %v %v = %v\n", num1, op, num2, result)
}
