package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Roman Структура римской цифры
type Roman struct {
	numeral string
	value   int
}

// RomanNumeral Массив для преобразования из строки в число
var RomanNumeral = []Roman{
	{"XC", 90}, {"C", 100}, {"XL", 40}, {"L", 50},
	{"IX", 9}, {"X", 10}, {"IV", 4}, {"V", 5},
	{"I", 1},
}

// RomanValue Массив для преобразования из числа в строку
var RomanValue = []Roman{
	{"C", 100}, {"XC", 90}, {"L", 50}, {"XL", 40},
	{"X", 10}, {"IX", 9}, {"V", 5}, {"IV", 4},
	{"I", 1},
}

func RomanNumberToInt(romanNumber string) (result int) {
	count := len(romanNumber)
	runes := []rune(romanNumber)

	i := 0
	for i < count {
		substring := string(runes[i:count])

		var matching Roman

		for _, glyph := range RomanNumeral {
			if strings.HasPrefix(substring, glyph.numeral) {
				matching = glyph
				break
			}
		}

		result += matching.value
		i += len(matching.numeral)
	}

	return result
}

// IntToRomanNumber converts a base 10 integer to its roman number equivalent
func IntToRomanNumber(number int) (result string) {
	for _, romanGlyph := range RomanValue {
		for number >= romanGlyph.value {
			result += romanGlyph.numeral
			number -= romanGlyph.value
		}
	}

	return result
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multi(a int, b int) int {
	return a * b
}

func divi(a int, b int) int {
	return a / b
}

func main() {
	var result int
	var isRoman bool

	fmt.Println("Калькулятор3000")
	fmt.Println("Напиши выражение одной строчкой с помощью двух операндов и одним оператором (+, -, /, *)")

	reader := bufio.NewReader(os.Stdin) // Считываем значения пользователя

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	elemets := strings.Fields(text) // Получаем элементы ввода

	// Проверяем ввод на правильность по кол-во элементов
	if len(elemets) > 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}
	if len(elemets) < 3 {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	// Преобразуем в int
	num1, er1 := strconv.Atoi(elemets[0])
	num2, er2 := strconv.Atoi(elemets[2])

	// Делаем выводы после преобразования. Либо выводим ошибку, либо идём дальше
	if (er1 == nil && er2 != nil) || (er1 != nil && er2 == nil) {
		fmt.Println("Вывод ошибки, так как числа должна быть целыми и быть одной системы счисления")
		return
	}
	if er1 != nil && er2 != nil {
		isRoman = true

		// Преобразуем римские цифры в int
		num1 = RomanNumberToInt(elemets[0])
		num2 = RomanNumberToInt(elemets[2])
	}

	if (num1 > 10 || 1 > num1) || (num2 > 10 || 1 > num2) {
		fmt.Println("Числа должны быть от 1 до 10 включительно")
		return
	}

	switch elemets[1] {
	case "+":
		result = sum(num1, num2)
	case "-":
		result = sub(num1, num2)
	case "/":
		result = divi(num1, num2)
	case "*":
		result = multi(num1, num2)
	default:
		println("Вывод ошибки, так как заданого знака не существует")
		return
	}

	if isRoman {
		if result < 1 {
			println("Вывод ошибки, тк в римской системе нет отрицательных чисел")
			return
		}
		println(IntToRomanNumber(result))

	} else {
		println(result)
	}
}
