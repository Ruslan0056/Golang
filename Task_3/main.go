package main

import (
	"fmt"
)

func main() {
	var text, keyword, message string
	var shifr uint
	fmt.Println("Выберите задание, где 1 - зашифровать, а 2 - расшифровать.")
	fmt.Scanln(&shifr)
	switch shifr {
	case 1:
		fmt.Println("Введите послание")
		fmt.Scanln(&text)
		fmt.Println("Введите ключ")
		fmt.Scanln(&keyword)
		keywordIndex := 0
		fmt.Println("Зашифрованное послание")
		for i := 0; i < len(text); i++ { // Итерирует каждый символ ASCII

			m := (text[i] + keyword[keywordIndex] - 'A') % 91

			if m < 65 {
				m += 65
			}
			message = string(m)
			fmt.Print(message)

			keywordIndex++
			keywordIndex %= len(keyword)

		}

		fmt.Println()

	case 2:
		fmt.Println("Введите зашифрованное послание")
		fmt.Scanln(&message)
		fmt.Println("Введите ключ")
		fmt.Scanln(&keyword)
		keywordIndex := 0
		fmt.Println("Расшифрованное послание")
		for i := 0; i < len(message); i++ { // Итерирует каждый символ ASCII

			m := (message[i] - keyword[keywordIndex] + 'A')

			if m < 65 {
				m += 26
			}

			text = string(m)
			fmt.Print(text)

			keywordIndex++
			keywordIndex %= len(keyword)

		}

		fmt.Println()
	default:
		fmt.Println("Неверный код задания")
	}

}
